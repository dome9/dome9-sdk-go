```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_notification"
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
	req.ChangeDetection.SNSSendingState = "Disabled"
	req.ChangeDetection.ExternalTicketCreatingState = "Disabled"
	req.ChangeDetection.AWSSecurityHubIntegrationState = "Disabled"
	req.ChangeDetection.WebhookIntegrationState = "Disabled"
	req.GCPSecurityCommandCenterIntegration.State = "Disabled"

    // Create CC Notification
	v, _, err := srv.Create(&req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create response type: %T\n Content %+v", v, v)

    // Get all CC Notifications
	resp, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", resp, resp)

    // Get specific CC Notification
	someNotification, _, err := srv.Get("SOME_ID")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", someNotification, someNotification)
    
    // Update specific CC Notification
	v, _, err := srv.Update("SOME_ID", &req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update response type: %T\n Content: %+v", v, v)
   
    // Delete CC Notification
    _, err := srv.Delete("SOME_ID")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Continuous Compliance Notification deleted")
}

```