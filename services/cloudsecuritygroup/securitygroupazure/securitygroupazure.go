package securitygroupazure

import (
	"fmt"
	"net/http"

	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup"
)

type CloudSecurityGroupRequest struct {
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	Region            string             `json:"region"`
	IsTamperProtected bool               `json:"isTamperProtected"`
	ResourceGroup     string             `json:"resourceGroup"`
	CloudAccountID    string             `json:"cloudAccountId"`
	Tags              []Tags             `json:"tags"`
	InboundServices   []InboundServices  `json:"inboundServices"`
	OutboundServices  []OutboundServices `json:"outboundServices"`
}

type Tags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//type Data struct {
//	Cidr string `json:"cidr"`
//	Note string `json:"note"`
//}

type Scope struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

type InboundServices struct {
	Direction             string   `json:"direction"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Priority              int      `json:"priority"`
	Access                string   `json:"access"`
	Protocol              string   `json:"protocol"`
	SourcePortRanges      []string `json:"sourcePortRanges"`
	SourceScopes          []Scope  `json:"sourceScopes"`
	DestinationPortRanges []string `json:"destinationPortRanges"`
	DestinationScopes     []Scope  `json:"destinationScopes"`
	IsDefault             bool     `json:"isDefault"`
}

type OutboundServices struct {
	Direction             string   `json:"direction"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Priority              int      `json:"priority"`
	Access                string   `json:"access"`
	Protocol              string   `json:"protocol"`
	SourcePortRanges      []string `json:"sourcePortRanges"`
	SourceScopes          []Scope  `json:"sourceScopes"`
	DestinationPortRanges []string `json:"destinationPortRanges"`
	DestinationScopes     []Scope  `json:"destinationScopes"`
	IsDefault             bool     `json:"isDefault"`
}

type CloudSecurityGroupResponse struct {
	ID                      string                      `json:"id"`
	ExternalSecurityGroupID string                      `json:"externalSecurityGroupId"`
	AccountID               int                         `json:"accountId"`
	CloudAccountName        string                      `json:"cloudAccountName"`
	LastUpdatedByDome9      bool                        `json:"lastUpdatedByDome9"`
	Error                   Error                       `json:"error"`
	CloudAccountID          string                      `json:"cloudAccountId"`
	Name                    string                      `json:"name"`
	Description             string                      `json:"description"`
	Region                  string                      `json:"region"`
	ResourceGroup           string                      `json:"resourceGroup"`
	InboundServices         []CloudSecurityGroupService `json:"inboundServices"`
	OutboundServices        []CloudSecurityGroupService `json:"outboundServices"`
	IsTamperProtected       bool                        `json:"isTamperProtected"`
	Tags                    []interface{}               `json:"tags"`
}

type CloudSecurityGroupService struct {
	Direction             string   `json:"direction"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Priority              int      `json:"priority"`
	Access                string   `json:"access"`
	Protocol              string   `json:"protocol"`
	SourcePortRanges      []string `json:"sourcePortRanges"`
	SourceScopes          []Scope  `json:"sourceScopes"`
	DestinationPortRanges []string `json:"destinationPortRanges"`
	DestinationScopes     []Scope  `json:"destinationScopes"`
	IsDefault             bool     `json:"isDefault"`
}

type Error struct {
	Action       string `json:"action"`
	ErrorMessage string `json:"errorMessage"`
}


func (service *Service) GetAll() (*[]CloudSecurityGroupResponse, *http.Response, error) {
	v := new([]CloudSecurityGroupResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudsecuritygroup.RESTfulPathAWS, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(id string) (*CloudSecurityGroupResponse, *http.Response, error) {
	v := new(CloudSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudsecuritygroup.RESTfulPathAzure, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *CloudSecurityGroupRequest) (*CloudSecurityGroupResponse, *http.Response, error) {
	v := new(CloudSecurityGroupResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudsecuritygroup.RESTfulPathAzure, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(id string, body *CloudSecurityGroupRequest) (*CloudSecurityGroupResponse, *http.Response, error) {
	v := new(CloudSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudsecuritygroup.RESTfulPathAzure, id)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudsecuritygroup.RESTfulPathAzure, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
