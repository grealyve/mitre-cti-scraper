package models

import "time"

type APTFeedConfigDB struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

// APTFeedData is a struct to hold the data from the APT Intrusion Set.
type APTFeedData struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	SpecVersion string `json:"spec_version"`
	Objects     []struct {
		Modified           time.Time `json:"modified"`
		Name               string    `json:"name"`
		Description        string    `json:"description"`
		Aliases            []string  `json:"aliases" gorm:"type:text"`
		XMitreDeprecated   bool      `json:"x_mitre_deprecated"`
		XMitreVersion      string    `json:"x_mitre_version"`
		XMitreContributors []string  `json:"x_mitre_contributors" gorm:"type:text"`
		Type               string    `json:"type"`
		ID                 string    `json:"id" gorm:"primary_key"`
		Created            time.Time `json:"created"`
		CreatedByRef       string    `json:"created_by_ref"`
		Revoked            bool      `json:"revoked"`
		ExternalReferences []struct {
			SourceName  string `json:"source_name"`
			URL         string `json:"url"`
			ExternalID  string `json:"external_id"`
			Description string `json:"description" gorm:"type:text"`
		} `json:"external_references" gorm:"type:jsonb;default:'[]';not null"`
		ObjectMarkingRefs       []string `json:"object_marking_refs" gorm:"type:text"`
		XMitreDomains           []string `json:"x_mitre_domains" gorm:"type:text"`
		XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
		XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
	} `json:"objects" gorm:"type:jsonb;default:'[]';not null"`
}

// APTFeedAttackPattern is a struct to hold the data from the APT Technics.
type APTFeedTechnic struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	SpecVersion string `json:"spec_version"`
	Objects []struct {
		XMitrePlatforms    []string  `json:"x_mitre_platforms" gorm:"type:text"`
		XMitreDomains      []string  `json:"x_mitre_domains" gorm:"type:text"`
		ObjectMarkingRefs  []string  `json:"object_marking_refs" gorm:"type:text"`
		Type               string    `json:"type"`
		ID                 string    `json:"id" gorm:"primary_key"`
		Created            time.Time `json:"created"`
		XMitreVersion      string    `json:"x_mitre_version"`
		ExternalReferences []struct {
			SourceName  string `json:"source_name"`
			URL         string `json:"url"`
			ExternalID  string `json:"external_id"`
			Description string `json:"description" gorm:"type:text"`
		} `json:"external_references" gorm:"type:jsonb;default:'[]';not null"`
		XMitreDeprecated bool      `json:"x_mitre_deprecated"`
		Revoked          bool      `json:"revoked"`
		Description      string    `json:"description"`
		Modified         time.Time `json:"modified"`
		CreatedByRef     string    `json:"created_by_ref"`
		Name             string    `json:"name"`
		XMitreDetection  string    `json:"x_mitre_detection"`
		KillChainPhases  []struct {
			KillChainName string `json:"kill_chain_name"`
			PhaseName     string `json:"phase_name"`
		} `json:"kill_chain_phases"`
		XMitreIsSubtechnique    bool     `json:"x_mitre_is_subtechnique"`
		XMitreTacticType        []string `json:"x_mitre_tactic_type" gorm:"type:text"`
		XMitreAttackSpecVersion string   `json:"x_mitre_attack_spec_version"`
		XMitreModifiedByRef     string   `json:"x_mitre_modified_by_ref"`
	} `json:"objects" gorm:"type:jsonb;default:'[]';not null"`
}

// APTFeedTactic is a struct to hold the data from the APT Tactic.
type APTFeedTactic struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	SpecVersion string `json:"spec_version"`
	Objects     []struct {
		XMitreDomains      []string  `json:"x_mitre_domains" gorm:"type:text"`
		ObjectMarkingRefs  []string  `json:"object_marking_refs" gorm:"type:text"`
		ID                 string    `json:"id" gorm:"primary_key"`
		Type               string    `json:"type"`
		Created            time.Time `json:"created"`
		CreatedByRef       string    `json:"created_by_ref"`
		ExternalReferences []struct {
			SourceName  string `json:"source_name"`
			URL         string `json:"url"`
			ExternalID  string `json:"external_id"`
			Description string `json:"description" gorm:"type:text"`
		} `json:"external_references" gorm:"type:jsonb;default:'[]';not null"`
		Modified                time.Time `json:"modified"`
		Name                    string    `json:"name"`
		Description             string    `json:"description"`
		XMitreVersion           string    `json:"x_mitre_version"`
		XMitreAttackSpecVersion string    `json:"x_mitre_attack_spec_version"`
		XMitreModifiedByRef     string    `json:"x_mitre_modified_by_ref"`
		XMitreShortname         string    `json:"x_mitre_shortname"`
	} `json:"objects" gorm:"type:jsonb;default:'[]';not null"`
}

// APTFeedRelationship is a struct to hold the data from the APT Relationship.
type APTFeedRelationship struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	SpecVersion string `json:"spec_version"`
	Objects     []struct {
		ObjectMarkingRefs  []string  `json:"object_marking_refs" gorm:"type:text"`
		Type               string    `json:"type"`
		ID                 string    `json:"id" gorm:"primary_key"`
		Created            time.Time `json:"created"`
		XMitreVersion      string    `json:"x_mitre_version"`
		ExternalReferences []struct {
			SourceName  string `json:"source_name"`
			URL         string `json:"url"`
			ExternalID  string `json:"external_id"`
			Description string `json:"description" gorm:"type:text"`
		} `json:"external_references" gorm:"type:jsonb;default:'[]';not null"`
		XMitreDeprecated        bool      `json:"x_mitre_deprecated"`
		Revoked                 bool      `json:"revoked"`
		Description             string    `json:"description"`
		Modified                time.Time `json:"modified"`
		CreatedByRef            string    `json:"created_by_ref"`
		RelationshipType        string    `json:"relationship_type"`
		SourceRef               string    `json:"source_ref"`
		TargetRef               string    `json:"target_ref"`
		XMitreAttackSpecVersion string    `json:"x_mitre_attack_spec_version"`
		XMitreModifiedByRef     string    `json:"x_mitre_modified_by_ref"`
	} `json:"objects" gorm:"type:jsonb;default:'[]';not null"`
}

// APTFeedMitigation is a struct to hold the data from the APT Mitigation.
type APTFeedMitigation struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	SpecVersion string `json:"spec_version"`
	Objects     []struct {
		XMitreDomains      []string  `json:"x_mitre_domains" gorm:"type:text"`
		ObjectMarkingRefs  []string  `json:"object_marking_refs" gorm:"type:text"`
		ID                 string    `json:"id" `
		Type               string    `json:"type"`
		Created            time.Time `json:"created"`
		CreatedByRef       string    `json:"created_by_ref"`
		ExternalReferences []struct {
			SourceName  string `json:"source_name"`
			URL         string `json:"url"`
			ExternalID  string `json:"external_id"`
			Description string `json:"description" gorm:"type:text"`
		} `json:"external_references" gorm:"type:jsonb;default:'[]';not null"`
		Modified            time.Time `json:"modified"`
		Name                string    `json:"name"`
		Description         string    `json:"description"`
		XMitreDeprecated    bool      `json:"x_mitre_deprecated"`
		XMitreVersion       string    `json:"x_mitre_version"`
		XMitreModifiedByRef string    `json:"x_mitre_modified_by_ref"`
	} `json:"objects" gorm:"type:jsonb;default:'[]';not null"`
}
