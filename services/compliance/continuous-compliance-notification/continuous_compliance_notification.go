package continuous_compliance_notification

import (
	"net/http"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/dome9/client"
)

const path = "Compliance/ContinuousComplianceNotification/"

type ScheduledData struct {
	CronExpression string   `json:"cronExpression"`
	Type           string   `json:"type"`
	Recipients     []string `json:"recipients"`
}

// EmailSendingState must be set to "Enabled" or "Disabled"
// ScheduledData has to be set to nil when EmailSendingState is Disabled
type ScheduledReport struct {
	EmailSendingState string         `json:"emailSendingState"`
	ScheduleData      *ScheduledData `json:"scheduleData"`
}

type EmailData struct {
	Recipients []string `json:"recipients"`
}

type EmailPerFindingData struct {
	Recipients               []string `json:"recipients"`
	NotificationOutputFormat string   `json:"notificationOutputFormat"`
}

type SnsData struct {
	SnsTopicArn     string `json:"snsTopicArn"`
	SnsOutputFormat string `json:"snsOutputFormat"`
}

type TicketingSystemData struct {
	SystemType         string `json:"systemType"`
	ShouldCloseTickets bool   `json:"shouldCloseTickets"`
	Domain             string `json:"domain"`
	User               string `json:"user"`
	Pass               string `json:"pass"`
	ProjectKey         string `json:"projectKey"`
	IssueType          string `json:"issueType"`
}

type AwsSecurityHubIntegration struct {
	ExternalAccountID string `json:"externalAccountId"`
	Region            string `json:"region"`
}

type WebhookData struct {
	URL        string `json:"url"`
	HTTPMethod string `json:"httpMethod"`
	AuthMethod string `json:"authMethod"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FormatType string `json:"formatType"`
}

// All the states must be set to "Enabled" or "Disabled"
// All the data or integration has to be set to nil when the corresponding status is Disabled
type ChangeDetection struct {
	EmailSendingState              string                     `json:"emailSendingState"`
	EmailPerFindingSendingState    string                     `json:"emailPerFindingSendingState"`
	SnsSendingState                string                     `json:"snsSendingState"`
	ExternalTicketCreatingState    string                     `json:"externalTicketCreatingState"`
	AwsSecurityHubIntegrationState string                     `json:"awsSecurityHubIntegrationState"`
	WebhookIntegrationState        string                     `json:"webhookIntegrationState"`
	EmailData                      *EmailData                 `json:"emailData"`
	EmailPerFindingData            *EmailPerFindingData       `json:"emailPerFindingData"`
	SnsData                        *SnsData                   `json:"snsData"`
	TicketingSystemData            *TicketingSystemData       `json:"ticketingSystemData"`
	AwsSecurityHubIntegration      *AwsSecurityHubIntegration `json:"awsSecurityHubIntegration"`
	WebhookData                    *WebhookData               `json:"webhookData"`
}

// State must be set to "Enabled" or "Disabled"
type GcpSecurityCommandCenterIntegration struct {
	State     string `json:"state"`
	ProjectID string `json:"projectId"`
	SourceID  string `json:"sourceId"`
}

type ContinuousComplianceNotificationPost struct {
	Name                                string                              `json:"name"`
	Description                         string                              `json:"description"`
	AlertsConsole                       bool                                `json:"alertsConsole"`
	ScheduledReport                     ScheduledReport                     `json:"scheduledReport"`
	ChangeDetection                     ChangeDetection                     `json:"changeDetection"`
	GcpSecurityCommandCenterIntegration GcpSecurityCommandCenterIntegration `json:"gcpSecurityCommandCenterIntegration"`
}

type ContinuousComplianceNotificationGet struct {
	ID                                  string                              `json:"id"`
	Name                                string                              `json:"name"`
	Description                         string                              `json:"description"`
	AlertsConsole                       bool                                `json:"alertsConsole"`
	ScheduledReport                     ScheduledReport                     `json:"scheduledReport"`
	ChangeDetection                     ChangeDetection                     `json:"changeDetection"`
	GcpSecurityCommandCenterIntegration GcpSecurityCommandCenterIntegration `json:"gcpSecurityCommandCenterIntegration"`
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) Create(body *ContinuousComplianceNotificationPost) (*ContinuousComplianceNotificationGet, *http.Response, error) {
	v := new(ContinuousComplianceNotificationGet)
	resp, err := service.client.NewRequestDo("POST", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() ([]ContinuousComplianceNotificationGet, *http.Response, error) {
	var v []ContinuousComplianceNotificationGet
	resp, err := service.client.NewRequestDo("GET", path, nil, nil, &v)

	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(id string, body *ContinuousComplianceNotificationPost) (*ContinuousComplianceNotificationGet, *http.Response, error) {
	v := new(ContinuousComplianceNotificationGet)
	resp, err := service.client.NewRequestDo("PUT", path+id, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	resp, err := service.client.NewRequestDo("DELETE", path+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
