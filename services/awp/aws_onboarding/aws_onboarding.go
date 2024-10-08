package awp_aws_onboarding

import (
	"fmt"
	"net/http"

	awp_onboarding "github.com/dome9/dome9-sdk-go/services/awp"
)

const (
	GetOnboardingDataPath = "workload/agentless/aws/terraform/onboarding"
)

type CreateAWPOnboardingRequestAws struct {
	CentralizedCloudAccountId  string                                   `json:"centralizedCloudAccountId"`
	CrossAccountRoleName       string                                   `json:"crossAccountRoleName"`
	CrossAccountRoleExternalId string                                   `json:"crossAccountRoleExternalId"`
	ScanMode                   string                                   `json:"scanMode,omitempty"`
	IsTerraform                bool                                     `json:"isTerraform"`
	AgentlessAccountSettings   *awp_onboarding.AgentlessAccountSettings `json:"agentlessAccountSettings"`
}

type AgentlessTerraformOnboardingDataResponseAws struct {
	Stage                                      string `json:"stage"`
	Region                                     string `json:"region"`
	CloudGuardBackendAccountId                 string `json:"cloudGuardBackendAccountId"`
	AgentlessBucketName                        string `json:"agentlessBucketName"`
	RemoteFunctionsPrefixKey                   string `json:"remoteFunctionsPrefixKey"`
	RemoteSnapshotsUtilsFunctionName           string `json:"remoteSnapshotsUtilsFunctionName"`
	RemoteSnapshotsUtilsFunctionRunTime        string `json:"remoteSnapshotsUtilsFunctionRunTime"`
	RemoteSnapshotsUtilsFunctionTimeOut        int    `json:"remoteSnapshotsUtilsFunctionTimeOut"`
	AwpClientSideSecurityGroupName             string `json:"awpClientSideSecurityGroupName"`
	RemoteSnapshotsUtilsFunctionS3PreSignedUrl string `json:"remoteSnapshotsUtilsFunctionCodePreSigneUrl"`
}

func (service *Service) CreateAWPOnboarding(id string, req CreateAWPOnboardingRequestAws, queryParams awp_onboarding.CreateOptions) (*http.Response, error) {
	pathPostfix := awp_onboarding.EnablePostfix
	if req.ScanMode == awp_onboarding.ScanModeInAccountSub {
		pathPostfix = awp_onboarding.EnableSubPostfix
		req.ScanMode = ""
	} else if req.ScanMode == awp_onboarding.ScanModeInAccountHub {
		pathPostfix = awp_onboarding.EnableHubPostfix
		req.ScanMode = ""
	}

	path := fmt.Sprintf(awp_onboarding.OnboardingResourcePath, awp_onboarding.ProviderAWS, id)
	return awp_onboarding.CreateAWPOnboarding(service.Client, req, fmt.Sprintf("%s/%s", path, pathPostfix), queryParams)
}

func (service *Service) GetAWPOnboarding(id string) (*awp_onboarding.GetAWPOnboardingResponse, *http.Response, error) {
	return awp_onboarding.GetAWPOnboarding(service.Client, awp_onboarding.ProviderAWS, id)
}

func (service *Service) DeleteAWPOnboarding(id string, queryParams awp_onboarding.DeleteOptions) (*http.Response, error) {
	return awp_onboarding.DeleteAWPOnboarding(service.Client, awp_onboarding.ProviderAWS, id, queryParams)
}

func (service *Service) UpdateAWPSettings(id string, scan_mode string, req awp_onboarding.AgentlessAccountSettings) (*http.Response, error) {
	pathPostfix := awp_onboarding.UpdatePostfix
	if scan_mode == awp_onboarding.ScanModeInAccountHub {
		pathPostfix = awp_onboarding.UpdateHubPostfix
	}

	path := fmt.Sprintf(awp_onboarding.OnboardingResourcePath, awp_onboarding.ProviderAWS, id)

	return awp_onboarding.UpdateAWPSettings(service.Client, fmt.Sprintf("%s/%s", path, pathPostfix), req)
}

func (service *Service) GetOnboardingData() (*AgentlessTerraformOnboardingDataResponseAws, *http.Response, error) {
	v := new(AgentlessTerraformOnboardingDataResponseAws)
	resp, err := service.Client.NewRequestDoRetry("GET", GetOnboardingDataPath, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
