package aws_org

import (
	_ "encoding/json"
	"time"
)

type CloudCredentialsType string

const (
	UserBased CloudCredentialsType = "UserBased"
	RoleBased CloudCredentialsType = "RoleBased"
)

type OnboardingPermissionRequest struct {
	RoleArn string               `json:"roleArn" validate:"required,roleArn"`
	Secret  string               `json:"secret" validate:"required,secret"`
	ApiKey  *string              `json:"apiKey,omitempty"`
	Type    CloudCredentialsType `json:"type" validate:"required,oneof=UserBased RoleBased"`
}

type ValidateStackSetArnRequest struct {
	OnboardingPermissionRequest
	StackSetArn string `json:"stackSetArn" validate:"required,stackSetArn"`
}

type OnboardingRequest struct {
	ValidateStackSetArnRequest
	AwsOrganizationName *string `json:"awsOrganizationName,omitempty"`
	EnableStackModify   bool    `json:"enableStackModify" validate:"required"`
}

type OnboardingUpdateRequest struct {
	AwsOrganizationName *string `json:"awsOrganizationName,omitempty"`
	EnableStackModify   bool    `json:"enableStackModify"`
}

type UpdateConfigurationRequest struct {
	OrganizationRootOuId *string                        `json:"organizationRootOuId" validate:"required"`
	MappingStrategy      MappingStrategyType            `json:"mappingStrategy" validate:"required"`
	PostureManagement    PostureManagementConfiguration `json:"postureManagement" validate:"required"`
}

type MappingStrategyType string
type OnboardingMode string

const (
	Flat   MappingStrategyType = "Flat"
	Clone  MappingStrategyType = "Clone"
	Read   OnboardingMode      = "Read"
	Manage OnboardingMode      = "Manage"
)

type PostureManagementConfiguration struct {
	RulesetsIds    []int64        `json:"rulesetsIds"`
	OnboardingMode OnboardingMode `json:"onboardingMode"`
}

type UpdateStackSetArnRequest struct {
	StackSetArn string `json:"stackSetArn" validate:"required,stackSetArn"`
}

type OrganizationOnboardingConfigurationBase struct {
	OrganizationRootOuId string                         `json:"organizationRootOuId,omitempty"`
	MappingStrategy      MappingStrategyType            `json:"mappingStrategy"`
	PostureManagement    PostureManagementConfiguration `json:"postureManagement"`
}

type AwsOrganizationOnboardingConfiguration struct {
	OrganizationOnboardingConfigurationBase
}

type OnboardingCftBase struct {
	ExternalId string `json:"externalId"`
	Content    string `json:"content"`
}

type ManagementCftConfiguration struct {
	OnboardingCftBase
	ManagementCftUrl      string `json:"managementCftUrl"`
	IsManagementOnboarded bool   `json:"isManagementOnboarded"`
}

type OnboardingMemberCft struct {
	OnboardingCftBase
	OnboardingCftUrl string `json:"onboardingCftUrl"`
}

type OrganizationManagementViewModel struct {
	Id                            string                                 `json:"id"`
	AccountId                     int64                                  `json:"accountId"`
	ExternalOrganizationId        string                                 `json:"externalOrganizationId"`
	ExternalManagementAccountId   string                                 `json:"externalManagementAccountId"`
	ManagementAccountStackId      string                                 `json:"managementAccountStackId"`
	ManagementAccountStackRegion  string                                 `json:"managementAccountStackRegion"`
	OnboardingConfiguration       AwsOrganizationOnboardingConfiguration `json:"onboardingConfiguration"`
	UserId                        int                                    `json:"userId"`
	EnableStackModify             bool                                   `json:"enableStackModify"`
	StackSetArn                   string                                 `json:"stackSetArn"`
	OrganizationName              string                                 `json:"organizationName"`
	UpdateTime                    time.Time                              `json:"updateTime"`
	CreationTime                  time.Time                              `json:"creationTime"`
	StackSetRegions               map[string]struct{}                    `json:"stackSetRegions"`
	StackSetOrganizationalUnitIds map[string]struct{}                    `json:"stackSetOrganizationalUnitIds"`
}
