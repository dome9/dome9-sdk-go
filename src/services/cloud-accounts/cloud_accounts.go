package cloud_accounts

import (
	"dome9"
	"dome9/client"
	"net/http"
	"time"
)

const (
	AwsResourceName   = "cloudaccounts"
	AzureResourceName = "AzureCloudAccount"
)

type AwsCredentials struct {
	ApiKey     string `json:"apikey"`
	Arn        string `json:"arn"`
	Secret     string `json:"secret"`
	IamUser    string `json:"iamUser"`
	Type       string `json:"type"`
	IsReadOnly bool   `json:"isReadOnly"`
}

type awsRegion struct {
	Region           string `json:"awsRegion"`
	Name             string `json:"name"`
	Hidden           bool   `json:"hidden"`
	NewGroupBehavior string `json:"newGroupBehavior"`
}

type awsNetSec struct {
	Regions []awsRegion `json:"regions"`
}

type AwsGetCloudAccountResponse struct {
	ID                     string         `json:"id"`
	Vendor                 string         `json:"vendor"`
	Name                   string         `json:"name"`
	ExternalAccountNumber  string         `json:"externalAccountNumber"`
	Error                  string         `json:"error"`
	IsFetchingSuspended    bool           `json:"isFetchingSuspended"`
	CreationDate           time.Time      `json:"creationDate"`
	Credentials            AwsCredentials `json:"credentials"`
	IamSafe                string         `json:"iamSafe"`
	NetSec                 awsNetSec      `json:"awsNetSec"`
	Magellan               bool           `json:"magellan"`
	FullProtection         bool           `json:"fullProtection"`
	AllowReadOnly          bool           `json:"allowReadOnly"`
	OrganizationalUnitID   string         `json:"organizationalUnitId"`
	OrganizationalUnitPath string         `json:"organizationalUnitPath"`
	OrganizationalUnitName string         `json:"organizationalUnitName"`
	LambdaScanner          bool           `json:"lambdaScanner"`
}

type AwsQueryParameters struct {
	ID string
}

// Required properties for onBoarding process
type AwsCreateRequest struct {
	Name              string `json:"name"`
	CustomCredentials struct {
		Arn    string `json:"arn"`
		Secret string `json:"secret"`
		Type   string `json:"type"`
	} `json:"credentials"`
	FullProtection bool `json:"fullProtection"`
	AllowReadOnly  bool `json:"allowReadOnly"`
}

type AzureCredentials struct {
	ClientID       string `json:"clientId"`
	ClientPassword string `json:"clientPassword"`
}

type AzureGetCloudAccountResponse struct {
	ID                     string           `json:"id"`
	Name                   string           `json:"name"`
	SubscriptionID         string           `json:"subscriptionId"`
	TenantID               string           `json:"tenantId"`
	Credentials            AzureCredentials `json:"credentials"`
	OperationMode          string           `json:"operationMode"`
	Error                  string           `json:"error"`
	CreationDate           time.Time        `json:"creationDate"`
	OrganizationalUnitID   string           `json:"organizationalUnitId"`
	OrganizationalUnitPath string           `json:"organizationalUnitPath"`
	OrganizationalUnitName string           `json:"organizationalUnitName"`
	Vendor                 string           `json:"vendor"`
}

// Required properties for onBoarding process
type AzureCreateRequest struct {
	Name           string           `json:"name"`
	SubscriptionID string           `json:"subscriptionId"`
	TenantID       string           `json:"tenantId"`
	Credentials    AzureCredentials `json:"credentials"`
	OperationMode  string           `json:"operationMode"`
}

type AzureQueryParameters struct {
	ID string
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) Create(resourceName string, body interface{}) (interface{}, *http.Response, error) {
	v := getResponse(resourceName)
	resp, err := service.client.NewRequestDo("POST", resourceName, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(resourceName string, options interface{}) (interface{}, *http.Response, error) {
	v := getResponse(resourceName)
	resp, err := service.client.NewRequestDo("GET", resourceName, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(resourceName string, options interface{}) (*http.Response, error) {
	resp, err := service.client.NewRequestDo("DELETE", resourceName, options, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getResponse(resourceName string) interface{} {
	switch resourceName {

	case AwsResourceName:
		return new(AwsGetCloudAccountResponse)

	case AzureResourceName:
		return new(AzureGetCloudAccountResponse)

	default:
		panic("Invalid resource name")
	}
}
