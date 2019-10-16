package assessment

import (
	"net/http"
)

const (
	path = "assessment/bundleV2"
)

type RunBundleRequest struct {
	ID                     int    `json:"id"`
	Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string `json:"externalCloudAccountId"`
	CloudAccountID         string `json:"cloudAccountId"`
}

type RunBundleResponse struct {
	Request          Request          `json:"request"`
	Tests            []Test           `json:"tests"`
	LocationMetadata LocationMetadata `json:"locationMetadata"`
	TestEntities     interface{}      `json:"testEntities"`
	DataSyncStatus   []DataSyncStatus `json:"dataSyncStatus"`
	AssessmentPassed bool             `json:"assessmentPassed"`
	HasErrors        bool             `json:"hasErrors"`
	ID               int              `json:"id"`
}

type Request struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Cft                    Cft    `json:"cft"`
	IsCft                  bool   `json:"isCft"`
	Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string `json:"externalCloudAccountId"`
	CloudAccountID         string `json:"cloudAccountId"`
	Region                 string `json:"region"`
	CloudNetwork           string `json:"cloudNetwork"`
	CloudAccountType       string `json:"cloudAccountType"`
	RequestID              string `json:"requestId"`
}

type Cft struct {
	RootName string `json:"rootName"`
	Params   []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"params"`
	Files []struct {
		Name     string `json:"name"`
		Template string `json:"template"`
	} `json:"files"`
}

type Test struct {
	Error             string         `json:"error"`
	TestedCount       int            `json:"testedCount"`
	RelevantCount     int            `json:"relevantCount"`
	NonComplyingCount int            `json:"nonComplyingCount"`
	ExclusionStats    ExclusionStats `json:"exclusionStats"`
	EntityResults     []EntityResult `json:"entityResults"`
	Rule              Rule           `json:"rule"`
	TestPassed        bool           `json:"testPassed"`
}

type ExclusionStats struct {
	TestedCount       int `json:"testedCount"`
	RelevantCount     int `json:"relevantCount"`
	NonComplyingCount int `json:"nonComplyingCount"`
}

type EntityResult struct {
	ValidationStatus string      `json:"validationStatus"`
	IsRelevant       bool        `json:"isRelevant"`
	IsValid          bool        `json:"isValid"`
	IsExcluded       bool        `json:"isExcluded"`
	ExclusionID      string      `json:"exclusionId"`
	RemediationID    string      `json:"remediationId"`
	Error            string      `json:"error"`
	TestObj          interface{} `json:"testObj"`
}

type Rule struct {
	Name          string `json:"name"`
	Severity      string `json:"severity"`
	Logic         string `json:"logic"`
	Description   string `json:"description"`
	Remediation   string `json:"remediation"`
	ComplianceTag string `json:"complianceTag"`
	Domain        string `json:"domain"`
	Priority      string `json:"priority"`
	ControlTitle  string `json:"controlTitle"`
	RuleID        string `json:"ruleId"`
	LogicHash     string `json:"logicHash"`
	IsDefault     bool   `json:"isDefault"`
}

type LocationMetadata struct {
	Account      Account      `json:"account"`
	Region       *Region      `json:"region"`
	CloudNetwork CloudNetwork `json:"cloudNetwork"`
}

type Account struct {
	Srl        string `json:"srl"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
}

type Region struct {
	Srl        string `json:"srl"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
}

type CloudNetwork struct {
	Srl        string `json:"srl"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
}

type DataSyncStatus struct {
	EntityType                   string                         `json:"entityType"`
	RecentlySuccessfulSync       bool                           `json:"recentlySuccessfulSync"`
	GeneralFetchPermissionIssues bool                           `json:"generalFetchPermissionIssues"`
	EntitiesWithPermissionIssues []EntitiesWithPermissionIssues `json:"entitiesWithPermissionIssues"`
}

type EntitiesWithPermissionIssues struct {
	ExternalID            string `json:"externalId"`
	Name                  string `json:"name"`
	CloudVendorIdentifier string `json:"cloudVendorIdentifier"`
}

func (service *Service) RunBundle(body *RunBundleRequest) (*RunBundleResponse, *http.Response, error) {
	v := new(RunBundleResponse)
	resp, err := service.Client.NewRequestDo("POST", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
