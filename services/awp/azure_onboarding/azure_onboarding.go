package awp_azure_onboarding

import (
	"fmt"
	"net/http"

	awp_onboarding "github.com/dome9/dome9-sdk-go/services/awp"
)

const (
	GetOnboardingDataPath = "workload/agentless/azure/terraform"
)

type CreateAWPOnboardingRequestAzure struct {
	CentralizedCloudAccountId string                                   `json:"centralizedCloudAccountId"`
	ScanMode                  string                                   `json:"scanMode"`
	IsTerraform               bool                                     `json:"isTerraform"`
	ManagementGroupId         string                                   `json:"managementGroupId"`
	AgentlessAccountSettings  *awp_onboarding.AgentlessAccountSettings `json:"agentlessAccountSettings"`
}

type AgentlessTerraformOnboardingDataResponseAzure struct {
	Region                    string `json:"region"`
	AppClientId               string `json:"appClientId"`
	CloudAccountId            string `json:"CloudAccountId"`
	CentralizedCloudAccountId string `json:"CentralizedCloudAccountId"`
}

type GetAWPOnboardingDataRequestAzure struct {
	CentralizedId string `url:"centralizedId"`
}

func (service *Service) CreateAWPOnboarding(id string, req CreateAWPOnboardingRequestAzure, queryParams awp_onboarding.CreateOptions) (*http.Response, error) {
	pathPostfix := awp_onboarding.EnablePostfix
	if req.ScanMode == awp_onboarding.ScanModeInAccountSub {
		pathPostfix = awp_onboarding.EnableSubPostfix
	} else if req.ScanMode == awp_onboarding.ScanModeInAccountHub {
		pathPostfix = awp_onboarding.EnableHubPostfix
	}

	path := fmt.Sprintf(awp_onboarding.OnboardingResourcePath, awp_onboarding.ProviderAzure, id)
	return awp_onboarding.CreateAWPOnboarding(service.Client, req, fmt.Sprintf("%s/%s", path, pathPostfix), queryParams)
}

func (service *Service) GetAWPOnboarding(id string) (*awp_onboarding.GetAWPOnboardingResponse, *http.Response, error) {
	return awp_onboarding.GetAWPOnboarding(service.Client, awp_onboarding.ProviderAzure, id)
}

func (service *Service) DeleteAWPOnboarding(id string) (*http.Response, error) {
	return awp_onboarding.DeleteAWPOnboarding(service.Client, awp_onboarding.ProviderAzure, id, awp_onboarding.DeleteOptions{})
}

func (service *Service) UpdateAWPSettings(id string, scan_mode string, req awp_onboarding.AgentlessAccountSettings) (*http.Response, error) {
	pathPostfix := awp_onboarding.UpdatePostfix
	if scan_mode == awp_onboarding.ScanModeInAccountHub {
		pathPostfix = awp_onboarding.UpdateHubPostfix
	}

	path := fmt.Sprintf(awp_onboarding.OnboardingResourcePath, awp_onboarding.ProviderAzure, id)

	return awp_onboarding.UpdateAWPSettings(service.Client, fmt.Sprintf("%s/%s", path, pathPostfix), req)
}

func (service *Service) GetOnboardingData(id string, req GetAWPOnboardingDataRequestAzure) (*AgentlessTerraformOnboardingDataResponseAzure, *http.Response, error) {
	v := new(AgentlessTerraformOnboardingDataResponseAzure)
	path := fmt.Sprintf("%s/%s/onboarding", GetOnboardingDataPath, id)
	resp, err := service.Client.NewRequestDoRetry("GET", path, req, nil, v, nil)

	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
