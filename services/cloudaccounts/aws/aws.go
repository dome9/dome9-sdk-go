package aws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
)

type CloudAccountAWSRequest struct {
	Name                   string                     `json:"name"`
	Credentials            CloudAccountAWSCredentials `json:"credentials"`
	FullProtection         bool                       `json:"fullProtection,omitempty"`
	AllowReadOnly          bool                       `json:"allowReadOnly,omitempty"`
	OrganizationalUnitID   string                     `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                     `json:"organizationalUnitPath,omitempty"`
	OrganizationalUnitName string                     `json:"organizationalUnitName,omitempty"`
	LambdaScanner          bool                       `json:"lambdaScanner,omitempty"`
}

type CloudAccountAWSResponse struct {
	ID                     string                     `json:"id"`
	Vendor                 string                     `json:"vendor"`
	Name                   string                     `json:"name"`
	ExternalAccountNumber  string                     `json:"externalAccountNumber"`
	Error                  string                     `json:"error,omitempty"`
	IsFetchingSuspended    bool                       `json:"isFetchingSuspended"`
	CreationDate           time.Time                  `json:"creationDate"`
	Credentials            CloudAccountAWSCredentials `json:"credentials"`
	IamSafe                CloudAccountAWSIamSafe     `json:"iamSafe,omitempty"`
	NetSec                 CloudAccountAWSNetSec      `json:"netSec,omitempty"`
	Magellan               bool                       `json:"magellan"`
	FullProtection         bool                       `json:"fullProtection"`
	AllowReadOnly          bool                       `json:"allowReadOnly"`
	OrganizationalUnitID   string                     `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                     `json:"organizationalUnitPath"`
	OrganizationalUnitName string                     `json:"organizationalUnitName"`
	LambdaScanner          bool                       `json:"lambdaScanner"`
}

type CloudAccountUpdateRegionConfigRequest struct {
	CloudAccountID        string                      `json:"cloudAccountId,omitempty"`
	ExternalAccountNumber string                      `json:"externalAccountNumber,omitempty"`
	Data                  CloudAccountAWSNetSecRegion `json:"data,omitempty"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitId string `json:"organizationalUnitId,omitempty"`
}

type CloudAccountUpdateCredentialsRequest struct {
	CloudAccountID        string                     `json:"cloudAccountId,omitempty"`
	ExternalAccountNumber string                     `json:"externalAccountNumber,omitempty"`
	Data                  CloudAccountAWSCredentials `json:"data,omitempty"`
}

type CloudAccountUpdateNameRequest struct {
	CloudAccountID        string `json:"cloudAccountId,omitempty"`
	ExternalAccountNumber string `json:"externalAccountNumber,omitempty"`
	Data                  string `json:"data,omitempty"`
}

type CloudAccountAWSCredentials struct {
	ApiKey     string `json:"apikey,omitempty"`
	Arn        string `json:"arn,omitempty"`
	Secret     string `json:"secret,omitempty"`
	IamUser    string `json:"iamUser,omitempty"`
	Type       string `json:"type,omitempty"`
	IsReadOnly bool   `json:"isReadOnly,omitempty"`
}

type CloudAccountAWSNetSec struct {
	Regions []CloudAccountAWSNetSecRegion `json:"regions,omitempty"`
}

type CloudAccountAWSNetSecRegion struct {
	Region           string `json:"region,omitempty"`
	Name             string `json:"name,omitempty"`
	Hidden           bool   `json:"hidden,omitempty"`
	NewGroupBehavior string `json:"newGroupBehavior,omitempty"`
}

type CloudAccountAWSIamSafe struct {
	AwsGroupArn           string                     `json:"awsGroupArn,omitempty"`
	AwsPolicyArn          string                     `json:"awsPolicyArn,omitempty"`
	Mode                  string                     `json:"mode,omitempty"`
	State                 string                     `json:"state,omitempty"`
	ExcludedIamEntities   CloudAccountAWSIamEntities `json:"excludedIamEntities,omitempty"`
	RestrictedIamEntities CloudAccountAWSIamEntities `json:"restrictedIamEntities,omitempty"`
}

type CloudAccountAWSIamEntities struct {
	RolesArns []string `json:"rolesArns,omitempty"`
	UsersArns []string `json:"usersArns,omitempty"`
}

func (service *Service) Get(options interface{}) (*CloudAccountAWSResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}

	v := new(CloudAccountAWSResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAWS, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]CloudAccountAWSResponse, *http.Response, error) {
	v := new([]CloudAccountAWSResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAWS, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountAWSRequest) (*CloudAccountAWSResponse, *http.Response, error) {
	v := new(CloudAccountAWSResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathAWS, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(body CloudAccountUpdateNameRequest) (*CloudAccountAWSResponse, *http.Response, error) {
	v := new(CloudAccountAWSResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, cloudaccounts.RESTfulServicePathAWSName)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateRegionConfig(body CloudAccountUpdateRegionConfigRequest) (*CloudAccountAWSResponse, *http.Response, error) {
	v := new(CloudAccountAWSResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, cloudaccounts.RESTfulServicePathAWSRegionConfig)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// TODO: not implemented in TF provider due to bug https://dome9-security.atlassian.net/browse/DOME-12538
func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountAWSResponse, *http.Response, error) {
	v := new(CloudAccountAWSResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAWS, id, cloudaccounts.RESTfulServicePathAWSOrganizationalUnit)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateCredentials(body CloudAccountUpdateCredentialsRequest) (*CloudAccountAWSResponse, *http.Response, error) {
	v := new(CloudAccountAWSResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, cloudaccounts.RESTfulServicePathAWSCredentials)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
