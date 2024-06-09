package azure_org

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws_org"
)

// Define the necessary types

type CloudVendor string

const (
	CloudVendorAzure      CloudVendor = "azure"
	CloudVendorAzureChina CloudVendor = "azurechina"
	CloudVendorAzureGov   CloudVendor = "azuregov"
)

type Blades struct {
	Awp               AwpConfiguration        `json:"awp" validate:"required"`
	Serverless        ServerlessConfiguration `json:"serverless" validate:"required"`
	Cdr               CdrConfiguration        `json:"cdr" validate:"required"`
	PostureManagement PostureManagement       `json:"postureManagement" validate:"required"`
}

type AwpOnboardingMode string

const (
	AwpOnboardingModeSaas         AwpOnboardingMode = "saas"
	AwpOnboardingModeInAccount    AwpOnboardingMode = "inAccount"
	AwpOnboardingModeInAccountHub AwpOnboardingMode = "inAccountHub"
)

type BladeConfiguration struct {
	IsEnabled bool `json:"isEnabled"`
}

type AwpConfiguration struct {
	BladeConfiguration
	OnboardingMode            AwpOnboardingMode `json:"onboardingMode"`
	CentralizedSubscriptionId string            `json:"centralizedSubscriptionId,omitempty"`
	WithFunctionAppsScan      bool              `json:"withFunctionAppsScan"`
}

type ServerlessConfiguration struct {
	BladeConfiguration
}

type StorageAccount struct {
	StorageId string   `json:"storageId"`
	LogTypes  []string `json:"logTypes"`
}

type CdrConfiguration struct {
	BladeConfiguration
	Accounts []StorageAccount `json:"accounts"`
}

type PostureManagement struct {
	OnboardingMode aws_org.OnboardingMode `json:"onboardingMode"`
}

type OnboardingRequest struct {
	WorkflowId              string      `json:"workflowId,omitempty"`
	TenantId                string      `json:"tenantId" validate:"required"`
	ManagementGroupId       string      `json:"managementGroupId,omitempty"`
	OrganizationName        string      `json:"organizationName,omitempty"`
	AppRegistrationName     string      `json:"appRegistrationName,omitempty"`
	ClientId                string      `json:"clientId,omitempty"`
	ClientSecret            string      `json:"clientSecret,omitempty"`
	ActiveBlades            Blades      `json:"activeBlades" validate:"required"`
	Vendor                  CloudVendor `json:"vendor" validate:"required,oneof=azure azurechina azuregov"`
	UseCloudGuardManagedApp bool        `json:"useCloudGuardManagedApp"`
	IsAutoOnboarding        bool        `json:"isAutoOnboarding"`
}

type OrganizationManagementViewModel struct {
	Id                      string                                   `json:"id"`
	AccountId               int64                                    `json:"accountId"`
	UserId                  int                                      `json:"userId"`
	OrganizationName        string                                   `json:"organizationName"`
	TenantId                string                                   `json:"tenantId"`
	ManagementGroupId       string                                   `json:"managementGroupId"`
	AppRegistrationName     string                                   `json:"appRegistrationName"`
	OnboardingConfiguration AzureOrganizationOnboardingConfiguration `json:"onboardingConfiguration"`
	UpdateTime              string                                   `json:"updateTime"`
	CreationTime            string                                   `json:"creationTime"`
	IsAutoOnboarding        bool                                     `json:"isAutoOnboarding"`
}

type AzureOrganizationOnboardingConfiguration struct {
	AwpConfiguration        *AwpConfiguration        `json:"awpConfiguration,omitempty"`
	ServerlessConfiguration *ServerlessConfiguration `json:"serverlessConfiguration,omitempty"`
	CdrConfiguration        *CdrConfiguration        `json:"cdrConfiguration,omitempty"`
	IsAutoOnboarding        bool                     `json:"isAutoOnboarding"`
}
