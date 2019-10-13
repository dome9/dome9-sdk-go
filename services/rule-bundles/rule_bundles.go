package rule_bundles

import (
	"net/http"
	"time"
)

const (
	path = "CompliancePolicy/"
)

type RuleBundleRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Rules       []struct {
		Name          string `json:"name,omitempty"`
		Severity      string `json:"severity,omitempty"`
		Logic         string `json:"logic,omitempty"`
		Description   string `json:"description,omitempty"`
		Remediation   string `json:"remediation,omitempty"`
		ComplianceTag string `json:"complianceTag,omitempty"`
		Domain        string `json:"domain,omitempty"`
		Priority      string `json:"priority,omitempty"`
		ControlTitle  string `json:"controlTitle,omitempty"`
		RuleID        string `json:"ruleId,omitempty"`
		LogicHash     string `json:"logicHash,omitempty"`
		IsDefault     bool   `json:"isDefault,omitempty"`
	} `json:"rules,omitempty"`
	ID               int    `json:"id,omitempty"`
	HideInCompliance bool   `json:"hideInCompliance,omitempty"`
	MinFeatureTier   string `json:"minFeatureTier,omitempty"`
	CloudVendor      string `json:"cloudVendor,omitempty"`
	Language         string `json:"language,omitempty"`
}

type RuleBundleResponse struct {
	Rules []struct {
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
	} `json:"rules"`
	AccountID        int       `json:"accountId"`
	CreatedTime      time.Time `json:"createdTime"`
	UpdatedTime      time.Time `json:"updatedTime"`
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	IsTemplate       bool      `json:"isTemplate"`
	HideInCompliance bool      `json:"hideInCompliance"`
	MinFeatureTier   string    `json:"minFeatureTier"`
	Section          int       `json:"section"`
	TooltipText      string    `json:"tooltipText"`
	ShowBundle       bool      `json:"showBundle"`
	SystemBundle     bool      `json:"systemBundle"`
	CloudVendor      string    `json:"cloudVendor"`
	Version          int       `json:"version"`
	Language         string    `json:"language"`
	RulesCount       int       `json:"rulesCount"`
}

func (service *Service) Get(id string) (*RuleBundleResponse, *http.Response, error) {
	v := new(RuleBundleResponse)
	resp, err := service.Client.NewRequestDo("GET", path+id, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAccountRuleBundles() (*[]RuleBundleResponse, *http.Response, error) {
	v := new([]RuleBundleResponse)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *RuleBundleRequest) (*RuleBundleResponse, *http.Response, error) {
	v := new(RuleBundleResponse)
	resp, err := service.Client.NewRequestDo("POST", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(body *RuleBundleRequest) (*RuleBundleResponse, *http.Response, error) {
	// Rule bundle ID passed within the request body
	v := new(RuleBundleResponse)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	resp, err := service.Client.NewRequestDo("DELETE", path+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
