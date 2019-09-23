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
	req.Vendor = "Azure"
	req.OperationMode = "Read"

	// Must fill below
	req.SubscriptionID = "AZURE SUBSCRIPTION ID FOR ACCOUNT"
	req.TenantID = "AZURE TENANT ID"
	req.Credentials.ClientID = "AZURE ACCOUNT ID"
	req.Credentials.ClientPassword = "PASSWORD FOR ACCOUNT"

	v, _, err := srv.Create(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create response type: %T\n Content %+v\n", v, v)

	resp, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v\n", resp, resp)
}
