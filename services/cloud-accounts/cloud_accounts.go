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
	GCPResourceNamePath   = "GoogleCloudAccount/"
)

type AwsCredentials struct {
	// ApiKey and Arn cannot exist together because Dome9 allows only one. To avoid passing empty strings (""),
	// we use pointers and the request is being sent with <nil>
	ApiKey     *string `json:"apikey"`
	Arn        *string `json:"arn"`
	Secret     string  `json:"secret"`
	IamUser    string  `json:"iamUser"`
	Type       string  `json:"type"`
	IsReadOnly bool    `json:"isReadOnly"`
}

type AwsRegion struct {
	Region           string `json:"awsRegion"`
	Name             string `json:"name"`
	Hidden           bool   `json:"hidden"`
	NewGroupBehavior string `json:"newGroupBehavior"`
}

type AwsNetSec struct {
	Regions []AwsRegion `json:"regions"`
}

// refer to API type: CloudAccountIamSafe
type AwsIAMSafe struct {
	AwsGroupArn           string         `json:"awsGroupArn"`
	AwsPolicyArn          string         `json:"awsPolicyArn"`
	Mode                  string         `json:"mode"`
	State                 string         `json:"state"`
	ExcludedIamEntities   AWSIamEntities `json:"excludedIamEntities"`
	RestrictedIamEntities AWSIamEntities `json:"restrictedIamEntities"`
}

type AWSIamEntities struct {
	RolesArns []string `json:"rolesArns"`
	UsersArns []string `json:"usersArns"`
}

// refer to API type: CloudAccount
type AwsGetCloudAccountResponse struct {
	ID                     string         `json:"id"`
	Vendor                 string         `json:"vendor"`
	Name                   string         `json:"name"`
	ExternalAccountNumber  string         `json:"externalAccountNumber"`
	Error                  string         `json:"error"`
	IsFetchingSuspended    bool           `json:"isFetchingSuspended"`
	CreationDate           time.Time      `json:"creationDate"`
	Credentials            AwsCredentials `json:"credentials"`
	IamSafe                AwsIAMSafe     `json:"iamSafe"`
	NetSec                 AwsNetSec      `json:"awsNetSec"`
	Magellan               bool           `json:"magellan"`
	FullProtection         bool           `json:"fullProtection"`
	AllowReadOnly          bool           `json:"allowReadOnly"`
	OrganizationalUnitID   string         `json:"organizationalUnitId"`
	OrganizationalUnitPath string         `json:"organizationalUnitPath"`
	OrganizationalUnitName string         `json:"organizationalUnitName"`
	LambdaScanner          bool           `json:"lambdaScanner"`
}

// refer to API type: CloudAccount
type AwsCreateRequest struct {
	Vendor                 string         `json:"vendor"`
	Name                   string         `json:"name"`
	ExternalAccountNumber  string         `json:"externalAccountNumber"`
	Error                  *string        `json:"error"`
	IsFetchingSuspended    bool           `json:"isFetchingSuspended"`
	CreationDate           time.Time      `json:"creationDate"`
	Credentials            AwsCredentials `json:"credentials"`
	FullProtection         bool           `json:"fullProtection"`
	AllowReadOnly          bool           `json:"allowReadOnly"`
	OrganizationalUnitID   string         `json:"organizationalUnitId"`
	OrganizationalUnitPath string         `json:"organizationalUnitPath"`
	OrganizationalUnitName string         `json:"organizationalUnitName"`
	LambdaScanner          bool           `json:"lambdaScanner"`
}

type AzureCredentials struct {
	ClientID       string `json:"clientId"`
	ClientPassword string `json:"clientPassword"`
}

// refer to API type: AzureCloudAccount
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

// refer to API type: AzureCloudAccount
type AzureCreateRequest struct {
	Name           string           `json:"name"`
	SubscriptionID string           `json:"subscriptionId"`
	TenantID       string           `json:"tenantId"`
	Credentials    AzureCredentials `json:"credentials"`
	OperationMode  string           `json:"operationMode"`
}

// refer to API type: GoogleCloudAccountGet
type GCPCloudAccountGetResponse struct {
	ID                     string    `json:"id"`
	Name                   string    `json:"name"`
	ProjectID              string    `json:"projectId"`
	CreationDate           time.Time `json:"creationDate"`
	OrganizationalUnitID   *string   `json:"organizationalUnitId"`
	OrganizationalUnitPath string    `json:"organizationalUnitPath"`
	OrganizationalUnitName string    `json:"organizationalUnitName"`
	Gsuite                 *struct {
		GsuiteUser string `json:"gsuiteUser"`
		DomainName string `json:"domainName"`
	} `json:"gsuite"`
	Vendor string `json:"vendor"`
}

// refer to API type: GoogleCloudAccountPost
type GCPCreateRequest struct {
	Name                      string `json:"name"`
	ServiceAccountCredentials struct {
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
	} `json:"serviceAccountCredentials"`
	GsuiteUser string `json:"gsuiteUser"`
	DomainName string `json:"domainName"`
}

type QueryParameters struct {
	ID string
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) CreateAWSCloudAccount(body interface{}) (*AwsGetCloudAccountResponse, *http.Response, error) {
	v := new(AwsGetCloudAccountResponse)
	resp, err := service.client.NewRequestDo("POST", AwsResourceNamePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) CreateGCPCloudAccount(body interface{}) (*GCPCloudAccountGetResponse, *http.Response, error) {
	v := new(GCPCloudAccountGetResponse)
	resp, err := service.client.NewRequestDo("POST", GCPResourceNamePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) CreateAzureCloudAccount(body interface{}) (*AzureGetCloudAccountResponse, *http.Response, error) {
	v := new(AzureGetCloudAccountResponse)
	resp, err := service.client.NewRequestDo("POST", AzureResourceNamePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAWSCloudAccount(options interface{}) (*AwsGetCloudAccountResponse, *http.Response, error) {
	v := new(AwsGetCloudAccountResponse)
	resp, err := service.client.NewRequestDo("GET", AwsResourceNamePath, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetGCPCloudAccount(options interface{}) (*GCPCloudAccountGetResponse, *http.Response, error) {
	v := new(GCPCloudAccountGetResponse)
	resp, err := service.client.NewRequestDo("GET", GCPResourceNamePath, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAzureCloudAccount(options interface{}) (*AzureGetCloudAccountResponse, *http.Response, error) {
	v := new(AzureGetCloudAccountResponse)
	resp, err := service.client.NewRequestDo("GET", AzureResourceNamePath, options, nil, v)
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
