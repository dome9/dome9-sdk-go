package cloud_accounts

import (
	"net/http"
	"time"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/dome9/client"
)

const (
	AwsResourceNamePath   = "cloudaccounts/"
	AzureResourceNamePath = "AzureCloudAccount/"
	GcpResourceNamePath   = "GoogleCloudAccount/"
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

type Gsuite struct {
	GsuiteUser string `json:"gsuiteUser"`
	DomainName string `json:"domainName"`
}

type GcpCloudAccountGet struct {
	ID                     string    `json:"id"`
	Name                   string    `json:"name"`
	ProjectID              string    `json:"projectId"`
	CreationDate           time.Time `json:"creationDate"`
	OrganizationalUnitID   string    `json:"organizationalUnitId"`
	OrganizationalUnitPath string    `json:"organizationalUnitPath"`
	OrganizationalUnitName string    `json:"organizationalUnitName"`
	Gsuite                 *Gsuite   `json:"gsuite"`
	Vendor                 string    `json:"vendor"`
}

type GcpServiceAccountCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

type GcpCloudAccountPost struct {
	Name                      string                       `json:"name"`
	ServiceAccountCredentials GcpServiceAccountCredentials `json:"serviceAccountCredentials"`
	GsuiteUser                string                       `json:"gsuiteUser"`
	DomainName                string                       `json:"domainName"`
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) Create(resourceNamePath string, body interface{}) (interface{}, *http.Response, error) {
	v := getResponse(resourceNamePath)
	resp, err := service.client.NewRequestDo("POST", resourceNamePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(resourceNamePath string, id string) (interface{}, *http.Response, error) {
	v := getResponse(resourceNamePath)
	resp, err := service.client.NewRequestDo("GET", resourceNamePath+id, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(resourceNamePath string, id string) (*http.Response, error) {
	resp, err := service.client.NewRequestDo("DELETE", resourceNamePath+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getResponse(resourceNamePath string) interface{} {
	switch resourceNamePath {

	case AwsResourceNamePath:
		return new(AwsGetCloudAccountResponse)

	case AzureResourceNamePath:
		return new(AzureGetCloudAccountResponse)

	case GcpResourceNamePath:
		return new(GcpCloudAccountGet)

	default:
		panic("Invalid resource name")
	}
}
