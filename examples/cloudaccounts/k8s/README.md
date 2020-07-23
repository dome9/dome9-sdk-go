```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := k8s.New(config)
	var createAccountReqBody k8s.CloudAccountRequest
	
	createAccountReqBody.Name = "ACCOUNT NAME" //mandatory field
	//createAccountReqBody.OrganizationalUnitID = '11111111-2222-3333-4444-555555555555' // optional field, if not set the root OU will be set as default 

	v, _, err := srv.Create(createAccountReqBody)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n Content %+v\n", v, v)

	// Get specific k8s cloud account
	k8sCloudAccount, _, err := srv.Get("ACCOUNT UUID")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n Content: %+v\n", k8sCloudAccount, k8sCloudAccount)
	
	// update k8s cloud account name
	id := "THE ACCOUNT ID"
	desiredNewName := "new k8s cloud account name"
	updateNameResponse, _, err := srv.UpdateName(id,
		k8s.CloudAccountUpdateNameRequest{
			Name: desiredNewName})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n Content: %+v\n", updateNameResponse, updateNameResponse)

	// update organizational Id
	OrganizationalUnitID := "ORGANIZATIONAL_UNIT_UUID"
	id = "THE ACCOUNT ID"
	OrganizationalUnitIDResponse, _, err := srv.UpdateOrganizationalID(id,
		k8s.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitId: OrganizationalUnitID})
	if err != nil {
		panic(err)
	}

	fmt.Printf("response type: %T\n Content: %+v\n", OrganizationalUnitIDResponse, OrganizationalUnitIDResponse)
    
    // Delete GCP cloud account
    _, err = srv.Delete("SOME_ID")
    if err != nil {
        panic(err)
    }

    fmt.Printf("K8S cloud accout deleted")
}

```