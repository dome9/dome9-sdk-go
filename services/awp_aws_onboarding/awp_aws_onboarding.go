package awp_aws_onboarding

import (
	"fmt"
	"log"
	"net/http"
)

const (
	awpAWSGetOnboardingDataPath = "workload/agentless/aws/terraform/onboarding"
	cloudAccountsPath           = "cloudaccounts/"
)

type AgentlessAwsTerraformOnboardingDataResponse struct {
	Stage                               string `json:"stage"`
	Region                              string `json:"region"`
	CloudGuardBackendAccountId          string `json:"cloudGuardBackendAccountId"`
	AgentlessBucketName                 string `json:"agentlessBucketName"`
	RemoteFunctionsPrefixKey            string `json:"remoteFunctionsPrefixKey"`
	RemoteSnapshotsUtilsFunctionName    string `json:"remoteSnapshotsUtilsFunctionName"`
	RemoteSnapshotsUtilsFunctionRunTime string `json:"remoteSnapshotsUtilsFunctionRunTime"`
	RemoteSnapshotsUtilsFunctionTimeOut int    `json:"remoteSnapshotsUtilsFunctionTimeOut"`
	AwpClientSideSecurityGroupName      string `json:"awpClientSideSecurityGroupName"`
}

type CloudAccountResponse struct {
	ID                     string      `json:"id"`
	Vendor                 string      `json:"vendor"`
	Name                   string      `json:"name"`
	ExternalAccountNumber  string      `json:"externalAccountNumber"`
	Error                  interface{} `json:"error"`
	IsFetchingSuspended    bool        `json:"isFetchingSuspended"`
	CreationDate           string      `json:"creationDate"`
	Credentials            Credentials `json:"credentials"`
	IamSafe                interface{} `json:"iamSafe"`
	NetSec                 NetSec      `json:"netSec"`
	Magellan               bool        `json:"magellan"`
	FullProtection         bool        `json:"fullProtection"`
	AllowReadOnly          bool        `json:"allowReadOnly"`
	OrganizationId         string      `json:"organizationId"`
	OrganizationalUnitId   interface{} `json:"organizationalUnitId"`
	OrganizationalUnitPath string      `json:"organizationalUnitPath"`
	OrganizationalUnitName string      `json:"organizationalUnitName"`
	LambdaScanner          bool        `json:"lambdaScanner"`
	Serverless             Serverless  `json:"serverless"`
	OnboardingMode         string      `json:"onboardingMode"`
}

type Credentials struct {
	Apikey     interface{} `json:"apikey"`
	Arn        string      `json:"arn"`
	Secret     interface{} `json:"secret"`
	IamUser    interface{} `json:"iamUser"`
	Type       string      `json:"type"`
	IsReadOnly bool        `json:"isReadOnly"`
}

type NetSec struct {
	Regions []Region `json:"regions"`
}

type Region struct {
	Region           string `json:"region"`
	Name             string `json:"name"`
	Hidden           bool   `json:"hidden"`
	NewGroupBehavior string `json:"newGroupBehavior"`
}

type Serverless struct {
	CodeAnalyzerEnabled           bool `json:"codeAnalyzerEnabled"`
	CodeDependencyAnalyzerEnabled bool `json:"codeDependencyAnalyzerEnabled"`
}

func (service *Service) Get() (*AgentlessAwsTerraformOnboardingDataResponse, *http.Response, error) {
	v := new(AgentlessAwsTerraformOnboardingDataResponse)
	resp, err := service.Client.NewRequestDo("GET", awpAWSGetOnboardingDataPath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetCloudAccountId(externalAccountId string) (string, *http.Response, error) {
	path := fmt.Sprintf("%s%s", cloudAccountsPath, externalAccountId)
	respData := new(CloudAccountResponse)
	log.Printf("[DEBUG] GetCloudAccountId Path: %s", path)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, respData)
	if err != nil {
		return "", nil, err
	}
	return respData.ID, resp, nil
}
