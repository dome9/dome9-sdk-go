package aws_org

import (
	_ "encoding/json"
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
	StackSetArn string `json:"stackSetArn" validate:"required,stackSetArn"`
}
