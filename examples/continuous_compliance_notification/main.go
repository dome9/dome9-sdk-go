package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/compliance/continuous_compliance_notification"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := continuous_compliance_notification.New(config)
	var req continuous_compliance_notification.ContinuousComplianceNotificationRequest

	req.Name = "test-1"
	req.Description = "test description"
	req.AlertsConsole = true
	req.ScheduledReport.EmailSendingState = "Disabled"
	req.ChangeDetection.EmailSendingState = "Disabled"
	req.ChangeDetection.EmailPerFindingSendingState = "Disabled"
	req.ChangeDetection.SnsSendingState = "Disabled"
	req.ChangeDetection.ExternalTicketCreatingState = "Disabled"
	req.ChangeDetection.AwsSecurityHubIntegrationState = "Disabled"
	req.ChangeDetection.WebhookIntegrationState = "Disabled"
	req.GcpSecurityCommandCenterIntegration.State = "Disabled"

	v, _, err := srv.Create(&req)
	resp, _, _ := srv.GetAll()
	fmt.Printf("Create response type: %T\n Content %+v", v, v)
	fmt.Printf("Get response type: %T\n Content: %+v", resp, resp)

	if err != nil {
		fmt.Println(err)
	}

}
