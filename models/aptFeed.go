package models

import (
	"github.com/lib/pq"

	"time"
)

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
	Type              string    `json:"type"`
	ID                string    `json:"id" gorm:"primaryKey"`
	SpecVersion       string    `json:"spec_version"`
	APTFeedDataObject JSONBData `json:"objects" gorm:"type:jsonb;not null"`
}

type APTFeedDataObject struct {
	Modified                time.Time               `json:"modified"`
	Name                    string                  `json:"name"`
	Description             string                  `json:"description"`
	Aliases                 pq.StringArray          `json:"aliases" gorm:"type:text[]"`
	XMitreDeprecated        bool                    `json:"x_mitre_deprecated"`
	XMitreVersion           string                  `json:"x_mitre_version"`
	XMitreContributors      pq.StringArray          `json:"x_mitre_contributors" gorm:"type:text[]"`
	Type                    string                  `json:"type"`
	ID                      string                  `json:"id" gorm:"primary_key"`
	Created                 time.Time               `json:"created"`
	CreatedByRef            string                  `json:"created_by_ref"`
	Revoked                 bool                    `json:"revoked"`
	ExternalReferences      JSONBExternalReferences `json:"external_references" gorm:"type:jsonb;default:'[]'"`
	ObjectMarkingRefs       pq.StringArray          `json:"object_marking_refs" gorm:"type:text[]"`
	XMitreDomains           pq.StringArray          `json:"x_mitre_domains" gorm:"type:text[]"`
	XMitreAttackSpecVersion string                  `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string                  `json:"x_mitre_modified_by_ref"`
}

// APTFeedAttackPattern is a struct to hold the data from the APT Technics.
type APTFeedTechnic struct {
	Type                 string       `json:"type"`
	ID                   string       `json:"id" gorm:"primaryKey"`
	SpecVersion          string       `json:"spec_version"`
	APTFeedTechnicObject JSONBTechnic `json:"objects" gorm:"type:jsonb;not null"`
}

type APTFeedTechnicObject struct {
	XMitrePlatforms         pq.StringArray          `json:"x_mitre_platforms" gorm:"type:text[]"`
	XMitreDomains           pq.StringArray          `json:"x_mitre_domains" gorm:"type:text[]"`
	ObjectMarkingRefs       pq.StringArray          `json:"object_marking_refs" gorm:"type:text[]"`
	Type                    string                  `json:"type"`
	ID                      string                  `json:"id" gorm:"primaryKey"`
	Created                 time.Time               `json:"created"`
	XMitreVersion           string                  `json:"x_mitre_version"`
	ExternalReferences      JSONBExternalReferences `json:"external_references" gorm:"type:jsonb;default:'[]'"`
	XMitreDeprecated        bool                    `json:"x_mitre_deprecated"`
	Revoked                 bool                    `json:"revoked"`
	Description             string                  `json:"description"`
	Modified                time.Time               `json:"modified"`
	CreatedByRef            string                  `json:"created_by_ref"`
	Name                    string                  `json:"name"`
	XMitreDetection         string                  `json:"x_mitre_detection"`
	KillChainPhases         JSONBKillChainPhases    `json:"kill_chain_phases" gorm:"type:jsonb;default:'[]'"`
	XMitreIsSubtechnique    bool                    `json:"x_mitre_is_subtechnique"`
	XMitreTacticType        pq.StringArray          `json:"x_mitre_tactic_type" gorm:"type:text[]"`
	XMitreAttackSpecVersion string                  `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string                  `json:"x_mitre_modified_by_ref"`
}

type KillChainPhases struct {
	KillChainName string `json:"kill_chain_name"`
	PhaseName     string `json:"phase_name"`
}

// APTFeedTactic is a struct to hold the data from the APT Tactic.
type APTFeedTactic struct {
	Type                string      `json:"type"`
	ID                  string      `json:"id" gorm:"primaryKey"`
	SpecVersion         string      `json:"spec_version"`
	APTFeedTacticObject JSONBTactic `json:"objects" gorm:"type:jsonb;not null"`
}

type APTFeedTacticObject struct {
	XMitreDomains           pq.StringArray          `json:"x_mitre_domains" gorm:"type:text[]"`
	ObjectMarkingRefs       pq.StringArray          `json:"object_marking_refs" gorm:"type:text[]"`
	ID                      string                  `json:"id" gorm:"primary_key"`
	Type                    string                  `json:"type"`
	Created                 time.Time               `json:"created"`
	CreatedByRef            string                  `json:"created_by_ref"`
	ExternalReferences      JSONBExternalReferences `json:"external_references" gorm:"type:jsonb;default:'[]'"`
	Modified                time.Time               `json:"modified"`
	Name                    string                  `json:"name"`
	Description             string                  `json:"description"`
	XMitreVersion           string                  `json:"x_mitre_version"`
	XMitreAttackSpecVersion string                  `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string                  `json:"x_mitre_modified_by_ref"`
	XMitreShortname         string                  `json:"x_mitre_shortname"`
}

// APTFeedRelationship is a struct to hold the data from the APT Relationship.
type APTFeedRelationship struct {
	Type                      string            `json:"type"`
	ID                        string            `json:"id" gorm:"primaryKey"`
	SpecVersion               string            `json:"spec_version"`
	APTFeedRelationshipObject JSONBRelationship `json:"objects" gorm:"type:jsonb;default:'[]'"`
}

type APTFeedRelationshipObject struct {
	ObjectMarkingRefs       pq.StringArray          `json:"object_marking_refs" gorm:"type:text[]"`
	Type                    string                  `json:"type"`
	ID                      string                  `json:"id" gorm:"primary_key"`
	Created                 time.Time               `json:"created"`
	XMitreVersion           string                  `json:"x_mitre_version"`
	ExternalReferences      JSONBExternalReferences `json:"external_references" gorm:"type:jsonb;default:'[]'"`
	Revoked                 bool                    `json:"revoked"`
	Description             string                  `json:"description"`
	Modified                time.Time               `json:"modified"`
	CreatedByRef            string                  `json:"created_by_ref"`
	RelationshipType        string                  `json:"relationship_type"`
	SourceRef               string                  `json:"source_ref"`
	TargetRef               string                  `json:"target_ref"`
	XMitreAttackSpecVersion string                  `json:"x_mitre_attack_spec_version"`
	XMitreModifiedByRef     string                  `json:"x_mitre_modified_by_ref"`
}

// APTFeedMitigation is a struct to hold the data from the APT Mitigation.
type APTFeedMitigation struct {
	Type                    string          `json:"type"`
	ID                      string          `json:"id"`
	SpecVersion             string          `json:"spec_version"`
	APTFeedMitigationObject JSONBMitigation `json:"objects" gorm:"type:jsonb;not null"`
}

type APTFeedMitigationObject struct {
	XMitreDomains       pq.StringArray          `json:"x_mitre_domains" gorm:"type:text[]"`
	ObjectMarkingRefs   pq.StringArray          `json:"object_marking_refs" gorm:"type:text[]"`
	ID                  string                  `json:"id" gorm:"primaryKey"`
	Type                string                  `json:"type"`
	Created             time.Time               `json:"created"`
	CreatedByRef        string                  `json:"created_by_ref"`
	ExternalReferences  JSONBExternalReferences `json:"external_references" gorm:"type:jsonb;default:'[]'"`
	Modified            time.Time               `json:"modified"`
	Name                string                  `json:"name"`
	Description         string                  `json:"description"`
	XMitreDeprecated    bool                    `json:"x_mitre_deprecated"`
	XMitreVersion       string                  `json:"x_mitre_version"`
	XMitreModifiedByRef string                  `json:"x_mitre_modified_by_ref"`
}

type ExternalReferences struct {
	SourceName  string `json:"source_name"`
	URL         string `json:"url"`
	ExternalID  string `json:"external_id,omitempty"`
	Description string `json:"description" gorm:"type:text"`
}
