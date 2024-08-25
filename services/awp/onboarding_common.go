package awp_onboarding

import (
	"fmt"
	"net/http"

	"github.com/dome9/dome9-sdk-go/dome9/client"
)

const (
	ProviderAWS   = "aws"
	ProviderAzure = "azure"
)

const (
	OnboardingResourcePath = "workload/agentless/%s/accounts/%s"
	EnablePostfix          = "enable"
	EnableSubPostfix       = "enableSubAccount"
	EnableHubPostfix       = "enableCentralizedAccount"
)

const (
	ScanModeInAccountSub = "inAccountSub"
	ScanModeInAccountHub = "inAccountHub"
)

type CreateOptions struct {
	ShouldCreatePolicy string `url:"shouldCreatePolicy"`
}

type DeleteOptions struct {
	ForceDelete string `url:"forceDelete"`
}

type AgentlessAccountSettings struct {
	DisabledRegions              []string          `json:"disabledRegions"`
	ScanMachineIntervalInHours   int               `json:"scanMachineIntervalInHours"`
	MaxConcurrenceScansPerRegion int               `json:"maxConcurrenceScansPerRegion"`
	SkipFunctionAppsScan         bool              `json:"skipFunctionAppsScan"`
	InAccountScannerVPC          string            `json:"inAccountScannerVPC"`
	CustomTags                   map[string]string `json:"customTags"`
}

type AgentlessCentralizedAccountSettings struct {
	InAccountScannerVPC string `json:"inAccountScannerVPC"`
}

type AccountIssues struct {
	Regions map[string]interface{}  `json:"regions"`
	Account *map[string]interface{} `json:"account"`
}

type GetAWPOnboardingResponse struct {
	AgentlessAccountSettings        *AgentlessAccountSettings `json:"agentlessAccountSettings"`
	MissingAwpPrivateNetworkRegions *[]string                 `json:"missingAwpPrivateNetworkRegions"`
	AccountIssues                   *AccountIssues            `json:"accountIssues"`
	CloudAccountId                  string                    `json:"cloudAccountId"`
	AgentlessProtectionEnabled      bool                      `json:"agentlessProtectionEnabled"`
	ScanMode                        string                    `json:"scanMode"`
	Provider                        string                    `json:"provider"`
	ShouldUpdate                    bool                      `json:"shouldUpdate"`
	IsOrgOnboarding                 bool                      `json:"isOrgOnboarding"`
	CentralizedCloudAccountId       string                    `json:"centralizedCloudAccountId"`
}

// Common functionality

func CreateAWPOnboarding(client *client.Client, req interface{}, path string, queryParams CreateOptions) (*http.Response, error) {
	resp, err := client.NewRequestDoRetry("POST", path, queryParams, req, nil, shouldRetry)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetAWPOnboarding(client *client.Client, cloudProvider string, id string) (*GetAWPOnboardingResponse, *http.Response, error) {
	v := new(GetAWPOnboardingResponse)
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)
	resp, err := client.NewRequestDoRetry("GET", path, nil, nil, v, shouldRetry)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func DeleteAWPOnboarding(client *client.Client, cloudProvider string, id string, queryParams DeleteOptions) (*http.Response, error) {
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)
	resp, err := client.NewRequestDoRetry("DELETE", path, queryParams, nil, nil, shouldRetry)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateAWPSettings(client *client.Client, cloudProvider string, id string, req AgentlessAccountSettings) (*http.Response, error) {
	// Construct the URL path
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)
	// Make a PATCH request with the JSON body
	resp, err := client.NewRequestDoRetry("PATCH", fmt.Sprintf("%s/settings", path), nil, req, nil, shouldRetry)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateAWPCentralizedSettings(client *client.Client, cloudProvider string, id string, req AgentlessCentralizedAccountSettings) (*http.Response, error) {
	// Construct the URL path
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)
	// Make a PATCH request with the JSON body
	resp, err := client.NewRequestDoRetry("PATCH", fmt.Sprintf("%s/settings", path), nil, req, nil, shouldRetry)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func shouldRetry(resp *http.Response) bool {
	return resp != nil && resp.StatusCode >= 400 && resp.StatusCode < 600
}
