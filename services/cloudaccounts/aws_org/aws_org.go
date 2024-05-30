package aws_org

import (
	_ "encoding/json"
)

// CloudCredentialsType enum
type CloudCredentialsType string

const (
	UserBased CloudCredentialsType = "UserBased"
	RoleBased CloudCredentialsType = "RoleBased"
)

type OnboardingPermissionRequest struct {
	// CloudGuard role arn from AWS.
	// AWS China accounts supported only in CloudGuard China DC.
	RoleArn string `json:"roleArn" validate:"required,roleArn"`

	// Also known as ExternalId from management-stack API.
	// this value should be the one that returns from: management-stack API.
	Secret string `json:"secret" validate:"required,secret"`

	// Should be null.
	// Needed only for 'UserBased' Type.
	ApiKey *string `json:"apiKey,omitempty"`

	// Default value is 'RoleBased'.
	// 'UserBased' is not supported.
	Type CloudCredentialsType `json:"type" validate:"required,oneof=UserBased RoleBased"`
}

type ValidateStackSetArnRequest struct {
	OnboardingPermissionRequest

	// The created StackSet ARN.
	StackSetArn string `json:"stackSetArn" validate:"required,stackSetArn"`
}

type OnboardingRequest struct {
	ValidateStackSetArnRequest

	// Not required.
	AwsOrganizationName *string `json:"awsOrganizationName,omitempty"`

	// Required. Default is false, it's for future use.
	EnableStackModify bool `json:"enableStackModify" validate:"required"`
}

type OnboardingUpdateRequest struct {
	// AWS organization name.
	AwsOrganizationName *string `json:"awsOrganizationName,omitempty"`

	// Required. Default is false, it's for future use.
	EnableStackModify bool `json:"enableStackModify"`
}

type UpdateConfigurationRequest struct {
	OrganizationRootOuId *string                        `json:"organizationRootOuId" validate:"required"`
	MappingStrategy      MappingStrategyType            `json:"mappingStrategy" validate:"required"`
	PostureManagement    PostureManagementConfiguration `json:"postureManagement" validate:"required"`
}

type MappingStrategyType string

const (
	Flat  MappingStrategyType = "Flat"
	Clone MappingStrategyType = "Clone"
)

type PostureManagementConfiguration struct {
	RulesetsIds    []int64        `json:"rulesetsIds"`
	OnboardingMode OnboardingMode `json:"onboardingMode"`
}

type OnboardingMode string

const (
	Read   OnboardingMode = "Read"
	Manage OnboardingMode = "Manage"
)

type UpdateStackSetArnRequest struct {
	// The created StackSet ARN.
	StackSetArn string `json:"stackSetArn" validate:"required,stackSetArn"`
}
