package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grealyve/mitre-cti-scraper/helpers"
	"github.com/grealyve/mitre-cti-scraper/models"
)

var (
	config = &models.APTFeedConfigDB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	intrusionEndpoints, technicEndpoints, tacticEndpoints, mitigationEndpoints, relationshipEndpoints []string
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

func GetAptFeed(ctx *gin.Context) {
	err := helpers.InitiliazeWebSocketConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Cannot initilaize WebSocket connection with Client. If you want to use just HTTP API set API_ONLY=true"})
		return
	}
	time.Sleep(1 * time.Second)
	defer helpers.CloseWSConnection()

	aptFeedInput := ctx.Query("aptFeed")

	//Connect to the database
	// _, err = database.ConnectAndMigrateDB(config)
	// if err != nil {
	// 	ctx.JSON(500, gin.H{"error": "Failed to connect to the database"})
	// 	helpers.SendMessageWS("APT Feed", "Failed to connect to the database", "error")
	// }

	err = getFeedURLs()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get feed URLs"})
		helpers.SendMessageWS("APT Feed", "Failed to get feed URLs", "error")
		helpers.SendMessageWS("", "chista_EXIT_chista", "info")
	}

	if aptFeedInput == "technics" {
		go GetAptFeedTechnic()
		time.Sleep(10 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Technics": APTFeedTechnicDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Technics %v", APTFeedTechnicDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "tactics" {
		go GetAptFeedTactics()
		time.Sleep(10 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Tactics": APTFeedTacticDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Tactics %v", APTFeedTacticDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "mitigations" {
		go GetAptFeedMitigations()
		time.Sleep(10 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Mitigations": APTFeedMitigationDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Mitigations %v", APTFeedMitigationDataSlice), "info")
		helpers.SendMessageWS("", "mitre_EXIT_mitre", "info")
	} else if aptFeedInput == "relationships" {
		go GetAptFeedRelationships()
		time.Sleep(10 * time.Second)
		ctx.JSON(200, gin.H{"All Mitre Relationships": APTFeedRelationshipDataSlice})
		helpers.SendMessageWS("APT Feed", fmt.Sprintf("All Mitre Relationships %v", APTFeedRelationshipDataSlice), "info")
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

func GetAptFeedTechnic() {
	for _, technicEndpoint := range technicEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + technicEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
		}
		APTFeedTechnicData := models.APTFeedTechnic{}

		err = json.Unmarshal(body, &APTFeedTechnicData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
		}

		APTFeedTechnicDataSlice = append(APTFeedTechnicDataSlice, APTFeedTechnicData)
	}
}

func GetAptFeedTactics() {
	for _, tacticEndpoint := range tacticEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + tacticEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
		}
		APTFeedTacticData := models.APTFeedTactic{}

		err = json.Unmarshal(body, &APTFeedTacticData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
		}

		APTFeedTacticDataSlice = append(APTFeedTacticDataSlice, APTFeedTacticData)
	}
}

func GetAptFeedMitigations() {
	for _, mitigationEndpoint := range mitigationEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + mitigationEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
		}
		APTFeedMitigationData := models.APTFeedMitigation{}

		err = json.Unmarshal(body, &APTFeedMitigationData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
		}

		APTFeedMitigationDataSlice = append(APTFeedMitigationDataSlice, APTFeedMitigationData)
	}
}

func GetAptFeedRelationships() {
	for _, relationshipEndpoint := range relationshipEndpoints {
		// Send the GET request
		response, err := http.Get(RAW_URL + relationshipEndpoint)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to send GET request", "error")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to read response body", "error")
		}
		APTFeedRelationshipData := models.APTFeedRelationship{}

		err = json.Unmarshal(body, &APTFeedRelationshipData)
		if err != nil {
			helpers.SendMessageWS("APT Feed", "Failed to parse JSON data", "error")
		}

		APTFeedRelationshipDataSlice = append(APTFeedRelationshipDataSlice, APTFeedRelationshipData)
	}
}
