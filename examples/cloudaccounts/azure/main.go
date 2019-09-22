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
	req.SubscriptionID = "Azure subscription id for account"
	req.TenantID = "Azure tenant id"
	req.Vendor = "Azure"
	req.Credentials.ClientID = "Azure account id"
	req.Credentials.ClientPassword = "Password for account"
	req.OperationMode = "Read"

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
