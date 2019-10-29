```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/gcp"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := gcp.New(config)
	var req gcp.CloudAccountRequest

	// The following fields can be extracted from GCP Security Reviewer json file
	req.Name = "ACCOUNT NAME"
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
	fmt.Printf("Response type: %T\n Content %+v\n", v, v)

	// Get all cloud accounts
	resp, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n Content: %+v\n", resp, resp)
	
	// Get specific accounts
	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: "SOME_ID"}
	gcpCloudAccount, _, err := srv.Get(&getCloudAccountQueryParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n Content: %+v\n", gcpCloudAccount, gcpCloudAccount)
	
	// update cloud account name
	id := "THE ACCOUNT ID"
	desiredNewName := "new cloud account name"
	updateNameResponse, _, err := srv.UpdateName(id,
		gcp.CloudAccountUpdateNameRequest{
			Name: desiredNewName})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n Content: %+v\n", updateNameResponse, updateNameResponse)

	// update account gsuite
	id = "THE ACCOUNT ID"
	updateAccountGSuiteResponse, _, err := srv.UpdateAccountGSuite(id,
		gcp.GSuite{
			GSuiteUser: "EMAIL ADDRESS",
			DomainName: "DOMAIN NAME",
		})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n Content: %+v\n", updateAccountGSuiteResponse, updateAccountGSuiteResponse)

	// update google cloud account credentials
	id = "THE ACCOUNT ID"
	var updateCredentialsReq gcp.CloudAccountUpdateCredentialsRequest

	updateCredentialsReq.Name = "ACCOUNT NAME"
	updateCredentialsReq.ServiceAccountCredentials.Type = "service_account"
	updateCredentialsReq.ServiceAccountCredentials.AuthURI = "https://accounts.google.com/o/oauth2/auth"
	updateCredentialsReq.ServiceAccountCredentials.TokenURI = "https://oauth2.googleapis.com/token"
	updateCredentialsReq.ServiceAccountCredentials.AuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
	updateCredentialsReq.ServiceAccountCredentials.ProjectID = "PROJECT ID"
	updateCredentialsReq.ServiceAccountCredentials.PrivateKeyID = "PRIVATE KEY ID"
	updateCredentialsReq.ServiceAccountCredentials.PrivateKey = "PRIVATE KEY"
	updateCredentialsReq.ServiceAccountCredentials.ClientEmail = "CLIENT EMAIL"
	updateCredentialsReq.ServiceAccountCredentials.ClientID = "CLIENT ID"
	updateAccountCredentials, _, err := srv.UpdateCredentials(id, updateCredentialsReq)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response type: %T\n Content: %+v\n", updateAccountCredentials, updateAccountCredentials)

	// update organizational Id
	OrganizationalUnitID := "ORGANIZATIONAL_UNIT_ID"
	id = "THE ACCOUNT ID"
	OrganizationalUnitIDResponse, _, err := srv.UpdateOrganizationalID(id,
		gcp.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitID: OrganizationalUnitID})
	if err != nil {
		panic(err)
	}

	fmt.Printf("response type: %T\n Content: %+v\n", OrganizationalUnitIDResponse, OrganizationalUnitIDResponse)
    
    // Delete GCP cloud account
    _, err := srv.Delete("SOME_ID")
    if err != nil {
        panic(err)
    }

    fmt.Printf("GCP cloud accout deleted")
}

```