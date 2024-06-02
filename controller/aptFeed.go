package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grealyve/mitre-cti-scraper/database"
	"github.com/grealyve/mitre-cti-scraper/helpers"
	"github.com/grealyve/mitre-cti-scraper/models"
	"gorm.io/gorm"
)

// TODO: Dataların yüklenmesinin bitmesini bekleyip databaseden data çekilebilir.

var (
	databaseConnection *gorm.DB
	config             = &models.APTFeedConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Password: "seferalgan",
		User:     "minibots",
		DBName:   "mitre_cti_db",
		SSLMode:  "disable",
	}

	intrusionEndpoints, technicEndpoints, tacticEndpoints, mitigationEndpoints, relationshipEndpoints []string
	APTFeedDataSlice                                                                                  []models.APTFeedData
	APTFeedTechnicDataSlice                                                                           []models.APTFeedTechnic
	APTFeedTacticDataSlice                                                                            []models.APTFeedTactic
	APTFeedMitigationDataSlice                                                                        []models.APTFeedMitigation
	APTFeedRelationshipDataSlice                                                                      []models.APTFeedRelationship
)

const (
	REPOSITORY_URL = "https://api.github.com/repos/mitre/cti/git/trees/master?recursive=1"
	RAW_URL        = "https://raw.githubusercontent.com/mitre/cti/master/"
	FILE           = "https://raw.githubusercontent.com/mitre/cti/master/"
)

func init() {
	//Connect to the database
	db, err := database.ConnectAndMigrateDB(config)
	databaseConnection = db
	if err != nil {
		fmt.Errorf("Failed to connect to the database")
		helpers.SendMessageWS("APT Feed", "Failed to connect to the database", "error")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	}

}

func GetAptFeed(ctx *gin.Context) {
	err := helpers.InitiliazeWebSocketConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Cannot initilaize WebSocket connection with Client. If you want to use just HTTP API set API_ONLY=true"})
		return
	}
	time.Sleep(1 * time.Second)
	defer helpers.CloseWSConnection()

	aptFeedInput := ctx.Query("aptFeed")

	err = getFeedURLs()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get feed URLs"})
		helpers.SendMessageWS("APT Feed", "Failed to get feed URLs", "error")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	}

	if aptFeedInput == "technics" {
		go GetAptFeedTechnic(databaseConnection, ctx)
		time.Sleep(60 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Technics": APTFeedTechnicDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Technics %v", APTFeedTechnicDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "tactics" {
		go GetAptFeedTactics(databaseConnection, ctx)
		time.Sleep(30 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Tactics": APTFeedTacticDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Tactics %v", APTFeedTacticDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "mitigations" {
		go GetAptFeedMitigations(databaseConnection, ctx)
		time.Sleep(20 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Mitigations": APTFeedMitigationDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Mitigations %v", APTFeedMitigationDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "relationships" {
		go GetAptFeedRelationships(databaseConnection, ctx)
		time.Sleep(20 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Relationships": APTFeedRelationshipDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Relationships %v", APTFeedRelationshipDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "intrusions" {
		go GetAptFeedData(databaseConnection, ctx)
		time.Sleep(20 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Intrusions": APTFeedDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Intrusions %v", APTFeedDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "endpoints" {
		ctx.JSON(200, gin.H{"Technic Endpoints": technicEndpoints,
			"Tactic Endpoints":       tacticEndpoints,
			"Mitigation Endpoints":   mitigationEndpoints,
			"Relationship Endpoints": relationshipEndpoints,
			"Intrusion Endpoints":    intrusionEndpoints})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("Technic Endpoints %v\nTactic Endpoints %v\nMitigation Endpoints %v\nRelationship Endpoints %v\nIntrusion Endpoints %v", technicEndpoints, tacticEndpoints, mitigationEndpoints, relationshipEndpoints, intrusionEndpoints), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else {
		ctx.JSON(500, gin.H{"error": "Invalid input"})
		helpers.SendMessageWS("APT Feed", "Invalid input", "error")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	}

}

// getFeedURLs is a function to get the feed URL endpoints from the MITRE repository.
func getFeedURLs() error {
	// Send the GET request
	response, err := http.Get(REPOSITORY_URL)
	if err != nil {
		return errors.New("Failed to send GET request " + err.Error())
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Failed to read response body " + err.Error())
	}

	// Parse the JSON data
	var mitreRepoData map[string]interface{}
	err = json.Unmarshal(body, &mitreRepoData)
	if err != nil {
		return errors.New("Failed to parse JSON data " + err.Error())
	}

	// Regular expressions for filtering the paths
	intrusionSetRegex, _ := regexp.Compile(`(?:mobile-attack|enterprise-attack|pre-attack|ics-attack)\/intrusion-set\/.+\.json`)
	technicsRegex, _ := regexp.Compile(`(?:mobile-attack|enterprise-attack|pre-attack|ics-attack)\/attack-pattern\/.+\.json`)
	tacticsRegex, _ := regexp.Compile(`(?:mobile-attack|enterprise-attack|pre-attack|ics-attack)\/x-mitre-tactic\/.+\.json`)
	mitigationsRegex, _ := regexp.Compile(`(?:mobile-attack|enterprise-attack|pre-attack|ics-attack)\/course-of-action\/.+\.json`)
	relationshipsRegex, _ := regexp.Compile(`(?:mobile-attack|enterprise-attack|pre-attack|ics-attack)\/relationship\/.+\.json`)

	// Grab the paths
	items := mitreRepoData["tree"].([]interface{})
	for _, item := range items {
		itemMap := item.(map[string]interface{})
		path := itemMap["path"].(string)

		switch {
		case technicsRegex.MatchString(path):
			technicEndpoints = append(technicEndpoints, path)
		case tacticsRegex.MatchString(path):
			tacticEndpoints = append(tacticEndpoints, path)
		case mitigationsRegex.MatchString(path):
			mitigationEndpoints = append(mitigationEndpoints, path)
		case relationshipsRegex.MatchString(path):
			relationshipEndpoints = append(relationshipEndpoints, path)
		case intrusionSetRegex.MatchString(path):
			intrusionEndpoints = append(intrusionEndpoints, path)
		}
	}

	return nil
}

func GetAptFeedData(db *gorm.DB, ctx *gin.Context) {
	for _, intrusionEndpoint := range intrusionEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + intrusionEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send GET request"})
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}
		APTFeedData := models.APTFeedData{}

		err = json.Unmarshal(body, &APTFeedData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON data"})
			return
		}

		for _, obj := range APTFeedData.APTFeedDataObject {
			err = database.InsertAPTFeedDataObject(db, obj)
			if err != nil {
				helpers.SendMessageWS("APT Feed", "Failed to insert object data into the database", "error")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert object data into the database"})
			}
		}

		// Insert the data into the database
		err = database.InsertAPTFeedData(db, APTFeedData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to insert data into the database", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
			return
		}

		APTFeedDataSlice = append(APTFeedDataSlice, APTFeedData)
	}
}

func GetAptFeedTechnic(db *gorm.DB, ctx *gin.Context) {
	for _, technicEndpoint := range technicEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + technicEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send GET request"})
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}
		APTFeedTechnicData := models.APTFeedTechnic{}

		err = json.Unmarshal(body, &APTFeedTechnicData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON data"})
			return
		}

		for _, obj := range APTFeedTechnicData.APTFeedTechnicObject {
			err = database.InsertAPTFeedTechnicObject(db, obj)
			if err != nil {
				helpers.SendMessageWS("APT Feed", "Failed to insert object data into the database APTFeedTechnicObject", "error")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert object data into the database APTFeedTechnicObject"})
			}
		}

		// Insert the data into the database
		err = database.InsertAPTFeedTechnicData(db, APTFeedTechnicData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to insert data into the database", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
			return
		}

		APTFeedTechnicDataSlice = append(APTFeedTechnicDataSlice, APTFeedTechnicData)
	}
}

func GetAptFeedTactics(db *gorm.DB, ctx *gin.Context) {
	for _, tacticEndpoint := range tacticEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + tacticEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send GET request"})
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}
		APTFeedTacticData := models.APTFeedTactic{}

		err = json.Unmarshal(body, &APTFeedTacticData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON data"})
			return
		}

		for _, obj := range APTFeedTacticData.APTFeedTacticObject {
			err = database.InsertAPTFeedTacticObject(db, obj)
			if err != nil {
				helpers.SendMessageWS("APT Feed", "Failed to insert object data into the database APTFeedTacticObject", "error")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert object data into the database APTFeedTacticObject"})
			}
		}

		// Insert the data into the database
		err = database.InsertAPTFeedTacticData(db, APTFeedTacticData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to insert data into the database", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
			return
		}

		APTFeedTacticDataSlice = append(APTFeedTacticDataSlice, APTFeedTacticData)
	}
}

func GetAptFeedMitigations(db *gorm.DB, ctx *gin.Context) {
	for _, mitigationEndpoint := range mitigationEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + mitigationEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send GET request"})
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}
		APTFeedMitigationData := models.APTFeedMitigation{}

		err = json.Unmarshal(body, &APTFeedMitigationData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON data"})
			return
		}

		for _, obj := range APTFeedMitigationData.APTFeedMitigationObject {
			err = database.InsertAPTFeedMitigationObject(db, obj)
			if err != nil {
				helpers.SendMessageWS("APT Feed", "Failed to insert object data into the database", "error")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert object data into the database"})
			}
		}

		// Insert the data into the database
		err = database.InsertAPTFeedMitigationData(db, APTFeedMitigationData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to insert data into the database", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		}

		APTFeedMitigationDataSlice = append(APTFeedMitigationDataSlice, APTFeedMitigationData)
	}
}

func GetAptFeedRelationships(db *gorm.DB, ctx *gin.Context) {
	for _, relationshipEndpoint := range relationshipEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + relationshipEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send GET request"})
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			helpers.SendMessageWS("APT Feed", "Received non-OK HTTP status", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Received non-OK HTTP status"})
			return
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}
		APTFeedRelationshipData := models.APTFeedRelationship{}

		err = json.Unmarshal(body, &APTFeedRelationshipData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON data"})
			return
		}

		for _, obj := range APTFeedRelationshipData.APTFeedRelationshipObject {
			err = database.InsertAPTFeedRelationshipObject(db, obj)
			if err != nil {
				helpers.SendMessageWS("APT Feed", "Failed to insert object data into the database APTFeedRelationshipObject", "error")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert object data into the database APTFeedRelationshipObject"})
			}
		}

		// Insert the data into the database
		err = database.InsertAPTFeedRelationshipData(db, APTFeedRelationshipData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to insert data into the database APTFeedRelationshipData", "error")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database APTFeedRelationshipData"})
		}

		APTFeedRelationshipDataSlice = append(APTFeedRelationshipDataSlice, APTFeedRelationshipData)
	}
}
