package continuous_compliance_notification

import (
	"net/http"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/dome9/client"
)

const path = "Compliance/ContinuousComplianceNotification/"

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

type ContinuousComplianceNotificationRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	AlertsConsole   bool   `json:"alertsConsole"`
	ScheduledReport struct {
		// EmailSendingState must be set to "Enabled" or "Disabled"
		EmailSendingState string `json:"emailSendingState"`
		// ScheduledData has to be set to nil when EmailSendingState is Disabled
		ScheduleData *struct {
			CronExpression string   `json:"cronExpression"`
			Type           string   `json:"type"`
			Recipients     []string `json:"recipients"`
		} `json:"scheduleData"`
	} `json:"scheduledReport"`
	// All the states must be set to "Enabled" or "Disabled"
	// All the data or integration has to be set to nil when the corresponding status is Disabled
	ChangeDetection struct {
		EmailSendingState              string `json:"emailSendingState"`
		EmailPerFindingSendingState    string `json:"emailPerFindingSendingState"`
		SnsSendingState                string `json:"snsSendingState"`
		ExternalTicketCreatingState    string `json:"externalTicketCreatingState"`
		AwsSecurityHubIntegrationState string `json:"awsSecurityHubIntegrationState"`
		WebhookIntegrationState        string `json:"webhookIntegrationState"`
		EmailData                      *struct {
			Recipients []string `json:"recipients"`
		} `json:"emailData"`
		EmailPerFindingData *struct {
			Recipients               []string `json:"recipients"`
			NotificationOutputFormat string   `json:"notificationOutputFormat"`
		} `json:"emailPerFindingData"`
		SnsData *struct {
			SnsTopicArn     string `json:"snsTopicArn"`
			SnsOutputFormat string `json:"snsOutputFormat"`
		} `json:"snsData"`
		TicketingSystemData *struct {
			SystemType         string `json:"systemType"`
			ShouldCloseTickets bool   `json:"shouldCloseTickets"`
			Domain             string `json:"domain"`
			User               string `json:"user"`
			Pass               string `json:"pass"`
			ProjectKey         string `json:"projectKey"`
			IssueType          string `json:"issueType"`
		} `json:"ticketingSystemData"`
		AwsSecurityHubIntegration *struct {
			ExternalAccountID string `json:"externalAccountId"`
			Region            string `json:"region"`
		} `json:"awsSecurityHubIntegration"`
		WebhookData *struct {
			URL        string `json:"url"`
			HTTPMethod string `json:"httpMethod"`
			AuthMethod string `json:"authMethod"`
			Username   string `json:"username"`
			Password   string `json:"password"`
			FormatType string `json:"formatType"`
		} `json:"webhookData"`
	} `json:"changeDetection"`
	// State must be set to "Enabled" or "Disabled"
	GcpSecurityCommandCenterIntegration struct {
		State     string `json:"state"`
		ProjectID *string `json:"projectId"`
		SourceID  *string `json:"sourceId"`
	} `json:"gcpSecurityCommandCenterIntegration"`
}

type ContinuousComplianceNotificationResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	AlertsConsole   bool   `json:"alertsConsole"`
	ScheduledReport struct {
		EmailSendingState string `json:"emailSendingState"`
		ScheduleData      *struct {
			CronExpression string   `json:"cronExpression"`
			Type           string   `json:"type"`
			Recipients     []string `json:"recipients"`
		} `json:"scheduleData"`
	} `json:"scheduledReport"`
	ChangeDetection struct {
		EmailSendingState              string `json:"emailSendingState"`
		EmailPerFindingSendingState    string `json:"emailPerFindingSendingState"`
		SnsSendingState                string `json:"snsSendingState"`
		ExternalTicketCreatingState    string `json:"externalTicketCreatingState"`
		AwsSecurityHubIntegrationState string `json:"awsSecurityHubIntegrationState"`
		WebhookIntegrationState        string `json:"webhookIntegrationState"`
		EmailData                      *struct {
			Recipients []string `json:"recipients"`
		} `json:"emailData"`
		EmailPerFindingData *struct {
			Recipients               []string `json:"recipients"`
			NotificationOutputFormat string   `json:"notificationOutputFormat"`
		} `json:"emailPerFindingData"`
		SnsData *struct {
			SnsTopicArn     string `json:"snsTopicArn"`
			SnsOutputFormat string `json:"snsOutputFormat"`
		} `json:"snsData"`
		TicketingSystemData *struct {
			SystemType         string `json:"systemType"`
			ShouldCloseTickets bool   `json:"shouldCloseTickets"`
			Domain             string `json:"domain"`
			User               string `json:"user"`
			Pass               string `json:"pass"`
			ProjectKey         string `json:"projectKey"`
			IssueType          string `json:"issueType"`
		} `json:"ticketingSystemData"`
		AwsSecurityHubIntegration *struct {
			ExternalAccountID string `json:"externalAccountId"`
			Region            string `json:"region"`
		} `json:"awsSecurityHubIntegration"`
		WebhookData *struct {
			URL        string `json:"url"`
			HTTPMethod string `json:"httpMethod"`
			AuthMethod string `json:"authMethod"`
			Username   string `json:"username"`
			Password   string `json:"password"`
			FormatType string `json:"formatType"`
		} `json:"webhookData"`
	} `json:"changeDetection"`
	GcpSecurityCommandCenterIntegration struct {
		State     string `json:"state"`
		ProjectID *string `json:"projectId"`
		SourceID  *string `json:"sourceId"`
	} `json:"gcpSecurityCommandCenterIntegration"`
}

func (service *Service) GetAll() ([]ContinuousComplianceNotificationResponse, *http.Response, error) {
	var v []ContinuousComplianceNotificationResponse
	resp, err := service.client.NewRequestDo("GET", path, nil, nil, &v)

	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *ContinuousComplianceNotificationRequest) (*ContinuousComplianceNotificationResponse, *http.Response, error) {
	v := new(ContinuousComplianceNotificationResponse)
	resp, err := service.client.NewRequestDo("POST", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(id string, body *ContinuousComplianceNotificationRequest) (*ContinuousComplianceNotificationResponse, *http.Response, error) {
	v := new(ContinuousComplianceNotificationResponse)
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
