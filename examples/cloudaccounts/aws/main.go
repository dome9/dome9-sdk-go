package main

import (
	"fmt"
	"os"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts/aws"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := aws.New(config)
	var req aws.CloudAccountRequest

	req.Name = "test AWS cloud account"
	req.Credentials.Arn = os.Getenv("ARN")
	req.Credentials.Secret = os.Getenv("SECRET")
	req.Credentials.Type = "RoleBased"

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
}
