package cloudaccount

import (
	"net/http"
	"time"
)

// AWSCloudAccountRequest and AWSCloudAccountResponse refer to API type: CloudAccount
type AWSCloudAccountRequest struct {
	Vendor                string   `json:"vendor"`
	Name                  string    `json:"name"`
	ExternalAccountNumber string    `json:"externalAccountNumber"`
	Error                 *string   `json:"error"`
	IsFetchingSuspended   bool      `json:"isFetchingSuspended"`
	CreationDate          time.Time `json:"creationDate"`
	Credentials           struct {
		ApiKey     *string `json:"apikey"`
		Arn        *string `json:"arn"`
		Secret     string  `json:"secret"`
		IamUser    string  `json:"iamUser"`
		Type       string  `json:"type"`
		IsReadOnly bool    `json:"isReadOnly"`
	} `json:"credentials"`
	FullProtection         bool   `json:"fullProtection"`
	AllowReadOnly          bool   `json:"allowReadOnly"`
	OrganizationalUnitID   string `json:"organizationalUnitId"`
	OrganizationalUnitPath string `json:"organizationalUnitPath"`
	OrganizationalUnitName string `json:"organizationalUnitName"`
	LambdaScanner          bool   `json:"lambdaScanner"`
}

type AWSCloudAccountResponse struct {
	ID                    string    `json:"id"`
	Vendor                string    `json:"vendor"`
	Name                  string    `json:"name"`
	ExternalAccountNumber string    `json:"externalAccountNumber"`
	Error                 *string   `json:"error"`
	IsFetchingSuspended   bool      `json:"isFetchingSuspended"`
	CreationDate          time.Time `json:"creationDate"`
	Credentials           struct {
		ApiKey     *string `json:"apikey"`
		Arn        *string `json:"arn"`
		Secret     *string `json:"secret"`
		IamUser    *string `json:"iamUser"`
		Type       string  `json:"type"`
		IsReadOnly bool    `json:"isReadOnly"`
	} `json:"credentials"`
	IamSafe *struct {
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
	} `json:"iamSafe"`
	NetSec *struct {
		Regions []struct {
			Region           string `json:"awsRegion"`
			Name             string `json:"name"`
			Hidden           bool   `json:"hidden"`
			NewGroupBehavior string `json:"newGroupBehavior"`
		} `json:"regions"`
	} `json:"netSec"`
	Magellan               bool    `json:"magellan"`
	FullProtection         bool    `json:"fullProtection"`
	AllowReadOnly          bool    `json:"allowReadOnly"`
	OrganizationalUnitID   *string `json:"organizationalUnitId"`
	OrganizationalUnitPath string  `json:"organizationalUnitPath"`
	OrganizationalUnitName string  `json:"organizationalUnitName"`
	LambdaScanner          bool    `json:"lambdaScanner"`
}

func (service *Service) GetCloudAccountAWS(options interface{}) (*AWSCloudAccountResponse, *http.Response, error) {
	v := new(AWSCloudAccountResponse)
	resp, err := service.client.NewRequestDo("GET", D9AwsResourceName, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) CreateCloudAccountAWS(body interface{}) (*AWSCloudAccountResponse, *http.Response, error) {
	v := new(AWSCloudAccountResponse)
	resp, err := service.client.NewRequestDo("POST", D9AwsResourceName, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
