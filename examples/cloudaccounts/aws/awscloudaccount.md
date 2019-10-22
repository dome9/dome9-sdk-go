```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := aws.New(config)
	var req aws.CloudAccountRequest

	req.Name = "test AWS cloud account"
	req.Credentials.Type = "RoleBased"

	// must fill below variables
	req.Credentials.Arn = "ARN"
	req.Credentials.Secret = "SECRET"

	// Create cloud account
	v, _, err := srv.Create(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create response type: %T\n Content %+v\n", v, v)

	// Get all cloud accounts
	cloudAccounts, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v\n", cloudAccounts, cloudAccounts)

	// Get specific cloud accounts
	cloudAccount, _, err := srv.Get("SOME_ID")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v\n", cloudAccount, cloudAccount)

	// update cloud account name
	desiredNewName := "test AWS update cloud account"
	updateNameResponse, _, err := srv.UpdateName(aws.CloudAccountUpdateNameRequest{
		CloudAccountID: v.ID,
		Data:           desiredNewName,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update name response type: %T\n Content: %+v\n", updateNameResponse, updateNameResponse)

	// Update Region Config
	desiredGroupBehavior := "FullManage"
	updateRegionConfigResponse, _, err := srv.UpdateRegionConfig(aws.CloudAccountUpdateRegionConfigRequest{
		CloudAccountID: v.ID,
		Data: aws.CloudAccountNetSecRegion{
			Region:           "us_east_1",
			NewGroupBehavior: desiredGroupBehavior,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update region config response type: %T\n Content: %+v\n", updateRegionConfigResponse, updateRegionConfigResponse)

	// Update Organizational Unit Id
	OrganizationalUnitID := "ORGANIZATIONAL_UNIT_ID"
	id := "THE ACCOUNT ID IN DOME9"
	updateOrganizationalIDResponse, _, err := srv.UpdateOrganizationalID(id,
		aws.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitId: OrganizationalUnitID},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("update Organizational ID response type: %T\n Content: %+v\n", updateOrganizationalIDResponse, updateOrganizationalIDResponse)

	// Update Credentials
	updateCredentialsResponse, _, err := srv.UpdateCredentials(aws.CloudAccountUpdateCredentialsRequest{
		CloudAccountID: v.ID,
		Data: aws.CloudAccountCredentials{
			Arn:    "ARN",
			Secret: "SECRET",
			Type:   "RoleBased",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update credentials response type: %T\n Content: %+v\n", updateCredentialsResponse, updateCredentialsResponse)
}

```