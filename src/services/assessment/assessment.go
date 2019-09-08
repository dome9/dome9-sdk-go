package assessment

import (
	"dome9"
	"dome9/client"
	"net/http"
)

const path  = "assessment/bundleV2"

type AssessmentBundleRequest struct {
	ID                     int    `json:"id"`
	Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string `json:"externalCloudAccountId"`
	CloudAccountID         string `json:"cloudAccountId"`
}

type params struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type file struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

type cft struct {
	RootName string   `json:"rootName"`
	Params   []params `json:"params"`
	Files    []file   `json:"files"`
}

type entityResults struct {
	ValidationStatus string      `json:"validationStatus"`
	IsRelevant       bool        `json:"isRelevant"`
	IsValid          bool        `json:"isValid"`
	IsExcluded       bool        `json:"isExcluded"`
	ExclusionID      string      `json:"exclusionId"`
	RemediationID    string      `json:"remediationId"`
	Error            string      `json:"error"`
	TestObj          interface{} `json:"testObj"`
}

type region struct {
	Srl        string `json:"srl"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
}

type request struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Cft                    cft    `json:"cft"`
	IsCft                  bool   `json:"isCft"`
	Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string `json:"externalCloudAccountId"`
	CloudAccountID         string `json:"cloudAccountId"`
	Region                 string `json:"region"`
	CloudNetwork           string `json:"cloudNetwork"`
	CloudAccountType       string `json:"cloudAccountType"`
	RequestID              string `json:"requestId"`
}

type exclusionStats struct {
	TestedCount       int `json:"testedCount"`
	RelevantCount     int `json:"relevantCount"`
	NonComplyingCount int `json:"nonComplyingCount"`
}

type rule struct {
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

type test struct {
	Error             string          `json:"error"`
	TestedCount       int             `json:"testedCount"`
	RelevantCount     int             `json:"relevantCount"`
	NonComplyingCount int             `json:"nonComplyingCount"`
	ExclusionStats    exclusionStats  `json:"exclusionStats"`
	EntityResults     []entityResults `json:"entityResults"`
	Rule              rule            `json:"rule"`
	TestPassed        bool            `json:"testPassed"`
}

type account struct {
	Srl        string `json:"srl"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
}

type cloudNetwork struct {
	Srl        string `json:"srl"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
}

type locationMetadata struct {
	Account      account      `json:"account"`
	Region       region       `json:"region"`
	CloudNetwork cloudNetwork `json:"cloudNetwork"`
}

type entitiesWithPermissionIssues struct {
	ExternalID            string `json:"externalId"`
	Name                  string `json:"name"`
	CloudVendorIdentifier string `json:"cloudVendorIdentifier"`
}

type dataSyncStatus struct {
	EntityType                   string                          `json:"entityType"`
	RecentlySuccessfulSync       bool                            `json:"recentlySuccessfulSync"`
	GeneralFetchPermissionIssues bool                            `json:"generalFetchPermissionIssues"`
	EntitiesWithPermissionIssues [] entitiesWithPermissionIssues `json:"entitiesWithPermissionIssues"`
}

type AssessmentResult struct {
	Request          request          `json:"request"`
	Tests            []test           `json:"tests"`
	LocationMetadata locationMetadata `json:"locationMetadata"`
	TestEntities     struct {
		NotSupported []interface{} `json:"notSupported"`
	} `json:"testEntities"`
	DataSyncStatus   []dataSyncStatus `json:"dataSyncStatus"`
	AssessmentPassed bool             `json:"assessmentPassed"`
	HasErrors        bool             `json:"hasErrors"`
	ID               int              `json:"id"`
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) RunAssessment(assessmentBundleRequest *AssessmentBundleRequest) (interface{}, *http.Response, error) {
	v := new(AssessmentResult)
	resp, err := service.client.NewRequestDo("POST", path, nil, assessmentBundleRequest, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
