package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts/azure"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := azure.New(config)
	var req azure.CloudAccountRequest

	req.Name = "test Azure cloud account"
	req.SubscriptionID = "SUBSCRIPTION_ID"
	req.TenantID = "TENANT_ID"
	req.Vendor = "Azure"
	req.Credentials.ClientID = "CLIENT_ID"
	req.Credentials.ClientPassword = "CLIENT_PASSWORD"
	req.OperationMode = "Read"

	v, _, err := srv.Create(req)
	resp, _, _ := srv.GetAll()
	fmt.Printf("Create response type: %T\n Content %+v\n", v, v)
	fmt.Printf("Get response type: %T\n Content: %+v\n", resp, resp)

	if err != nil {
		fmt.Println(err)
	}

}
