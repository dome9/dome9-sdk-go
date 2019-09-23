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

	// The following fields can be extracted from GCP Security Reviewer json file
	req.Name = "test GCP cloud account"
	req.ServiceAccountCredentials.Type = "service_account"
	req.ServiceAccountCredentials.AuthURI = "https://accounts.google.com/o/oauth2/auth"
	req.ServiceAccountCredentials.TokenURI = "https://oauth2.googleapis.com/token"
	req.ServiceAccountCredentials.AuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
	req.GsuiteUser = ""
	req.DomainName = ""

	// must fill below variables
	req.ServiceAccountCredentials.ProjectID = "PROJECT ID"
	req.ServiceAccountCredentials.PrivateKeyID = "PRIVATE KEY ID"
	req.ServiceAccountCredentials.PrivateKey = "PRIVATE KEY"
	req.ServiceAccountCredentials.ClientEmail = "CLIENT EMAIL"
	req.ServiceAccountCredentials.ClientID = "CLIENT ID"

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
