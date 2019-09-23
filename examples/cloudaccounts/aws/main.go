package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts/aws"
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

	// update cloud account name
	desiredNewName := "test AWS update cloud account"
	updateResponse, _, err := srv.UpdateName(aws.CloudAccountUpdateNameRequest{
		CloudAccountID: v.ID,
		Data:           desiredNewName,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update name response type: %T\n Content: %+v\n", updateResponse, updateResponse)

	// Update Region Config
	desiredGroupBehavior := "FullManage"
	updateRegionConfigResponse, _, err := srv.UpdateRegionConfig(aws.CloudAccountUpdateRegionConfigRequest{
		CloudAccountID: v.ID,
		Data: struct {
			Region           string `json:"region,omitempty"`
			Name             string `json:"name,omitempty"`
			Hidden           bool   `json:"hidden,omitempty"`
			NewGroupBehavior string `json:"newGroupBehavior,omitempty"`
		}{
			Region:           "us_east_1",
			NewGroupBehavior: desiredGroupBehavior,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("update region config response type: %T\n Content: %+v\n", updateRegionConfigResponse, updateRegionConfigResponse)

	// Update Organizational Unit Id
	OrganizationalUnitId := "ORGANIZATIONAL_UNIT_ID"
	updateOrganizationalIDResponse, _, err := srv.UpdateOrganizationalID(aws.CloudAccountUpdateOrganizationalIDRequest{
		OrganizationalUnitId: OrganizationalUnitId},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("update Organizational ID response type: %T\n Content: %+v\n", updateOrganizationalIDResponse, updateOrganizationalIDResponse)

	// Update Credentials
	updateCredentialsResponse, _, err := srv.UpdateCredentials(aws.CloudAccountUpdateCredentialsRequest{
		CloudAccountID: v.ID,
		Data: struct {
			Apikey     string `json:"apikey,omitempty"`
			Arn        string `json:"arn,omitempty"`
			Secret     string `json:"secret,omitempty"`
			IamUser    string `json:"iamUser,omitempty"`
			Type       string `json:"type,omitempty"`
			IsReadOnly bool   `json:"isReadOnly,omitempty"`
		}{
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
