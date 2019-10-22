```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
    "github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/azure"
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
	
	// Get all cloud accounts
	resp, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v\n", resp, resp)

	// Get specific account
	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: "SOME_ID"}
	azureCloudAccount, _, err := srv.Get(&getCloudAccountQueryParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v\n", azureCloudAccount, azureCloudAccount)
	
	// update cloud account name
	id := "THE ACCOUNT ID"
	desiredNewName := "new cloud account name"
	updateNameResponse, _, err := srv.UpdateName(id,
		azure.CloudAccountUpdateNameRequest{
			Name: desiredNewName})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update name response type: %T\n Content: %+v\n", updateNameResponse, updateNameResponse)

	// update operation mode
	id = "THE ACCOUNT ID IN DOME9"
	desiredMode := "Manage" // options: Manage, Read
	updateOperationMode, _, err := srv.UpdateOperationMode(id,
		azure.CloudAccountUpdateOperationModeRequest{
			OperationMode: desiredMode})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update operation mode response type: %T\n Content: %+v\n", updateOperationMode, updateOperationMode)

	// Update Organizational Unit Id
	id = "THE ACCOUNT ID IN DOME9"
	OrganizationalUnitId := "ORGANIZATIONAL UNIT ID"
	updateOrganizationalIDResponse, _, err := srv.UpdateOrganizationalID(id,
		azure.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitID: OrganizationalUnitId})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update Organizational ID response type: %T\n Content: %+v\n", updateOrganizationalIDResponse, updateOrganizationalIDResponse)

	// Update Credentials
	id = "THE ACCOUNT ID IN DOME9"
	updateCredentialsResponse, _, err := srv.UpdateCredentials(id,
		azure.CloudAccountUpdateCredentialsRequest{
			ApplicationID:  "APPLICATION ID",
			ApplicationKey: "APPLICATION KEY",
		})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update credentials response type: %T\n Content: %+v\n", updateCredentialsResponse, updateCredentialsResponse)

    // Delete Azure cloud account
    _, err := srv.Delete("SOME_ID")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Azure cloud accout deleted")
}

```