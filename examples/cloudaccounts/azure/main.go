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

	// update cloud account name
	desiredNewName := "test Azure update cloud account"
	updateNameResponse, _, err := srv.UpdateName(azure.CloudAccountUpdateNameRequest{Name: desiredNewName})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update name response type: %T\n Content: %+v\n", updateNameResponse, updateNameResponse)

	// update operation mode
	desiredMode := "Manage"
	updateOperationMode, _, err := srv.UpdateOperationMode(azure.CloudAccountUpdateOperationModeRequest{OperationMode: desiredMode})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update operation mode response type: %T\n Content: %+v\n", updateOperationMode, updateOperationMode)

	// Update Organizational Unit Id
	OrganizationalUnitId := "ORGANIZATIONAL_UNIT_ID"
	updateOrganizationalIDResponse, _, err := srv.UpdateOrganizationalID(azure.CloudAccountUpdateOrganizationalIDRequest{OrganizationalUnitID: OrganizationalUnitId})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update Organizational ID response type: %T\n Content: %+v\n", updateOrganizationalIDResponse, updateOrganizationalIDResponse)

	// Update Credentials
	updateCredentialsResponse, _, err := srv.UpdateCredentials(azure.CloudAccountUpdateCredentialsRequest{
		ApplicationID:  "APPLICATION ID",
		ApplicationKey: "APPLICATION KEY",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update credentials response type: %T\n Content: %+v\n", updateCredentialsResponse, updateCredentialsResponse)

}
