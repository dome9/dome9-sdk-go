package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := continuous_compliance_policy.New(config)
	var req continuous_compliance_policy.ContinuousCompliancePolicyRequest
	// Set cloud account parameters
	req.CloudAccountID = "CLOUD_ACCOUNT_ID"
	req.ExternalAccountID = "EXTERNAL_ACCOUNT_ID"
	req.CloudAccountType = "Aws"
	// Set Rule bundle ID
	req.BundleID = 86685
	// Set notification ID's
	req.NotificationIds = []string{"NOTIFICATION_IDS"}

	v, _, err := srv.Create(&req)
	resp, _, _ := srv.GetAll()
	fmt.Printf("Create response type: %T\n Content %+v", v, v)
	fmt.Printf("Get response type: %T\n Content: %+v", resp, resp)

	if err != nil {
		fmt.Println(err)
	}

}
