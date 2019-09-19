package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts/gcp"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := gcp.New(config)
	var req gcp.CloudAccountRequest

	req.Name = "test GCP cloud account"
	// The following fields can be extracted from GCP Security Reviewer json file
	req.ServiceAccountCredentials.Type = "service_account"
	req.ServiceAccountCredentials.ProjectID = "PROJECT_ID"
	req.ServiceAccountCredentials.PrivateKeyID = "PRIVATE_KEY_ID"
	req.ServiceAccountCredentials.PrivateKey = "PRIVATE_KEY"
	req.ServiceAccountCredentials.ClientEmail = "CLIENT_EMAIL"
	req.ServiceAccountCredentials.ClientID = "CLIENT_ID"
	req.ServiceAccountCredentials.AuthURI = "https://accounts.google.com/o/oauth2/auth"
	req.ServiceAccountCredentials.TokenURI = "https://oauth2.googleapis.com/token"
	req.ServiceAccountCredentials.AuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
	req.GsuiteUser = ""
	req.DomainName = ""

	v, _, err := srv.Create(req)
	resp, _, _ := srv.GetAll()
	fmt.Printf("Create response type: %T\n Content %+v\n", v, v)
	fmt.Printf("Get response type: %T\n Content: %+v\n", resp, resp)

	if err != nil {
		fmt.Println(err)
	}

}
