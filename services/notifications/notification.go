package notifications

import (
	"fmt"
	"net/http"
	"time"
)

const (
	RESTfulServicePathNotification = "notification"
)

// AssessmentFindingOrigin enum
type AssessmentFindingOrigin int

const (
	ComplianceEngine AssessmentFindingOrigin = iota
	Magellan
	MagellanAwsGuardDuty        = 191
	Serverless                  = 2
	AwsInspector                = 50
	ServerlessSecurityAnalyzer  = 51
	ExternalFindingSource       = 100
	Qualys                      = 101
	Tenable                     = 102
	AwsGuardDuty                = 103
	KubernetesImageScanning     = 104
	KubernetesRuntimeAssurance  = 105
	ContainersRuntimeProtection = 106
	WorkloadChangeMonitoring    = 107
	ImageAssurance              = 7
	SourceCodeAssurance         = 8
	InfrastructureAsCode        = 9
	CIEM                        = 10
	Incident                    = 11
)

// NotificationIntegrationSettingsModel struct
type NotificationIntegrationSettingsModel struct {
	ReportsIntegrationSettings            []ReportNotificationIntegrationSettings    `json:"reportsIntegrationSettings"`
	SingleNotificationIntegrationSettings []SingleNotificationIntegrationSettings    `json:"singleNotificationIntegrationSettings"`
	ScheduledIntegrationSettings          []ScheduledNotificationIntegrationSettings `json:"scheduledIntegrationSettings"`
}

// ReportNotificationIntegrationSettings struct
type ReportNotificationIntegrationSettings struct {
	BaseNotificationIntegrationSettings
}

// SingleNotificationIntegrationSettings struct
type SingleNotificationIntegrationSettings struct {
	BaseNotificationIntegrationSettings
	Payload string `json:"payload"`
}

// ScheduledNotificationIntegrationSettings struct
type ScheduledNotificationIntegrationSettings struct {
	BaseNotificationIntegrationSettings
	CronExpression string `json:"cronExpression" validate:"required,cron"`
}

// BaseNotificationIntegrationSettings struct
type BaseNotificationIntegrationSettings struct {
	IntegrationId string                       `json:"integrationId" validate:"required"`
	OutputType    int                          `json:"outputType"`
	Filter        ComplianceNotificationFilter `json:"filter"`
}

// NotificationTriggerType enum
type NotificationTriggerType int

const (
	Report NotificationTriggerType = iota
	Single
	Scheduled
)

// NotificationOutputType enum
type NotificationOutputType int

// ComplianceNotificationFilter struct
type ComplianceNotificationFilter struct{}

// BaseNotificationViewModel struct
type BaseNotificationViewModel struct {
	Name                 string                               `json:"name" validate:"required"`
	Description          string                               `json:"description"`
	AlertsConsole        bool                                 `json:"alertsConsole" default:"true"`
	SendOnEachOccurrence bool                                 `json:"sendOnEachOccurrence"`
	Origin               AssessmentFindingOrigin              `json:"origin" validate:"required"`
	IntegrationSettings  NotificationIntegrationSettingsModel `json:"integrationSettingsModel" validate:"required"`
}

// ResponseNotificationViewModel struct
type ResponseNotificationViewModel struct {
	BaseNotificationViewModel
	Id                   string                               `json:"id" validate:"required"`
	CreatedAt            time.Time                            `json:"createdAt" validate:"required"`
	Name                 string                               `json:"name" validate:"required"`
	Description          string                               `json:"description"`
	AlertsConsole        bool                                 `json:"alertsConsole" default:"true"`
	SendOnEachOccurrence bool                                 `json:"sendOnEachOccurrence"`
	Origin               string                               `json:"origin" validate:"required"`
	IntegrationSettings  NotificationIntegrationSettingsModel `json:"integrationSettingsModel" validate:"required"`
}

// Request models

// PutNotificationViewModel struct
type PutNotificationViewModel struct {
	BaseNotificationViewModel
	Id string `json:"id" validate:"required"`
}

// PostNotificationViewModel struct
type PostNotificationViewModel struct {
	BaseNotificationViewModel
}

// APIs

func (service *Service) Create(body PostNotificationViewModel) (*ResponseNotificationViewModel, *http.Response, error) {
	v := new(ResponseNotificationViewModel)
	resp, err := service.Client.NewRequestDo("POST", RESTfulServicePathNotification, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]ResponseNotificationViewModel, *http.Response, error) {
	v := new([]ResponseNotificationViewModel)
	resp, err := service.Client.NewRequestDo("GET", RESTfulServicePathNotification, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetById(id string) (*ResponseNotificationViewModel, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(ResponseNotificationViewModel)
	relativeURL := fmt.Sprintf("%s/%s", RESTfulServicePathNotification, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetByName(name string) (*ResponseNotificationViewModel, *http.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name parameter must be passed")
	}

	v := new(ResponseNotificationViewModel)
	relativeURL := fmt.Sprintf("%s?name=%s", RESTfulServicePathNotification, name)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(body PutNotificationViewModel) (*ResponseNotificationViewModel, *http.Response, error) {
	if body.Id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(ResponseNotificationViewModel)
	resp, err := service.Client.NewRequestDo("PUT", RESTfulServicePathNotification, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", RESTfulServicePathNotification, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
