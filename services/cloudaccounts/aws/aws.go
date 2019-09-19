package aws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts"
)

// AWSCloudAccountRequest and AWSCloudAccountResponse refer to API type: CloudAccounts
type CloudAccountRequest struct {
	Vendor                string    `json:"vendor"`
	Name                  string    `json:"name"`
	ExternalAccountNumber string    `json:"externalAccountNumber"`
	Error                 string    `json:"error,omitempty"`
	IsFetchingSuspended   bool      `json:"isFetchingSuspended"`
	CreationDate          time.Time `json:"creationDate"`
	Credentials           struct {
		ApiKey     string `json:"apikey,omitempty"`
		Arn        string `json:"arn,omitempty"`
		Secret     string `json:"secret"`
		IamUser    string `json:"iamUser"`
		Type       string `json:"type"`
		IsReadOnly bool   `json:"isReadOnly"`
	} `json:"credentials"`
	FullProtection         bool   `json:"fullProtection"`
	AllowReadOnly          bool   `json:"allowReadOnly"`
	OrganizationalUnitID   string `json:"organizationalUnitId"`
	OrganizationalUnitPath string `json:"organizationalUnitPath"`
	OrganizationalUnitName string `json:"organizationalUnitName"`
	LambdaScanner          bool   `json:"lambdaScanner"`
}

type CloudAccountResponse struct {
	ID                    string    `json:"id"`
	Vendor                string    `json:"vendor"`
	Name                  string    `json:"name"`
	ExternalAccountNumber string    `json:"externalAccountNumber"`
	Error                 string    `json:"error,omitempty"`
	IsFetchingSuspended   bool      `json:"isFetchingSuspended"`
	CreationDate          time.Time `json:"creationDate"`
	Credentials           struct {
		ApiKey     string `json:"apikey,omitempty"`
		Arn        string `json:"arn,omitempty"`
		Secret     string `json:"secret,omitempty"`
		IamUser    string `json:"iamUser,omitempty"`
		Type       string `json:"type"`
		IsReadOnly bool   `json:"isReadOnly"`
	} `json:"credentials"`
	IamSafe struct {
		AwsGroupArn         string `json:"awsGroupArn"`
		AwsPolicyArn        string `json:"awsPolicyArn"`
		Mode                string `json:"mode"`
		State               string `json:"state"`
		ExcludedIamEntities struct {
			RolesArns []string `json:"rolesArns"`
			UsersArns []string `json:"usersArns"`
		} `json:"excludedIamEntities"`
		RestrictedIamEntities struct {
			RolesArns []string `json:"rolesArns"`
			UsersArns []string `json:"usersArns"`
		} `json:"restrictedIamEntities"`
	} `json:"iamSafe,omitempty"`
	NetSec struct {
		Regions []struct {
			Region           string `json:"awsRegion"`
			Name             string `json:"name"`
			Hidden           bool   `json:"hidden"`
			NewGroupBehavior string `json:"newGroupBehavior"`
		} `json:"regions"`
	} `json:"netSec,omitempty"`
	Magellan               bool   `json:"magellan"`
	FullProtection         bool   `json:"fullProtection"`
	AllowReadOnly          bool   `json:"allowReadOnly"`
	OrganizationalUnitID   string `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string `json:"organizationalUnitPath"`
	OrganizationalUnitName string `json:"organizationalUnitName"`
	LambdaScanner          bool   `json:"lambdaScanner"`
}

func (service *Service) Get(options interface{}) (*CloudAccountResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAWS, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAWS, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathAWS, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeAddress := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeAddress, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
