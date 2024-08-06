package awp_onboarding

import (
	"fmt"
	"net/http"
	"time"

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

const (
	maxRetries    = 3
	retryInterval = time.Second * 5
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
	CustomTags                   map[string]string `json:"customTags"`
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

func retryableRequest(f func() (*http.Response, error)) (*http.Response, error) {
	var resp *http.Response
	var err error
	var lastStatusCode int

	for i := 0; i < maxRetries; i++ {
		resp, err = f()
		if err == nil {
			return resp, nil
		}

		if resp != nil {
			lastStatusCode = resp.StatusCode
		}

		if resp != nil && (resp.StatusCode >= 400 && resp.StatusCode < 500) {
			time.Sleep(retryInterval)
		} else {
			break
		}
	}

	return nil, fmt.Errorf("request failed after %d attempts with status code %d: %w", maxRetries, lastStatusCode, err)
}

func CreateAWPOnboarding(client *client.Client, req interface{}, path string, queryParams CreateOptions) (*http.Response, error) {
	f := func() (*http.Response, error) {
		return client.NewRequestDo("POST", path, queryParams, req, nil)
	}

	return retryableRequest(f)
}

func GetAWPOnboarding(client *client.Client, cloudProvider string, id string) (*GetAWPOnboardingResponse, *http.Response, error) {
	v := new(GetAWPOnboardingResponse)
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)

	f := func() (*http.Response, error) {
		return client.NewRequestDo("GET", path, nil, nil, v)
	}

	resp, err := retryableRequest(f)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func DeleteAWPOnboarding(client *client.Client, cloudProvider string, id string, queryParams DeleteOptions) (*http.Response, error) {
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)

	f := func() (*http.Response, error) {
		return client.NewRequestDo("DELETE", path, queryParams, nil, nil)
	}

	return retryableRequest(f)
}

func UpdateAWPSettings(client *client.Client, cloudProvider string, id string, req AgentlessAccountSettings) (*http.Response, error) {
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)

	f := func() (*http.Response, error) {
		return client.NewRequestDo("PATCH", fmt.Sprintf("%s/settings", path), nil, req, nil)
	}

	return retryableRequest(f)
}
