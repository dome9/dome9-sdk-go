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
	req.CloudAccountType = "Aws"

	// Set Rule bundle ID
	desiredBundleID := 86685
	req.BundleID = desiredBundleID

	// must fill below variables
	req.CloudAccountID = "CLOUD ACCOUNT ID"
	req.ExternalAccountID = "EXTERNAL ACCOUNT ID"
	req.NotificationIds = []string{"NOTIFICATION ID"}

	v, _, err := srv.Create(&req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create response type: %T\n Content %+v", v, v)

	resp, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v", resp, resp)
}
