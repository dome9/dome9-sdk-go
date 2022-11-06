```go
package main

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/imageassurance/imageassurance_policy"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("Access-ID", "Secret-Key", "https://api.dome9.com/v2/")
	srv := imageassurance_policy.New(config)

	request := imageassurance_policy.ImageAssurancePolicyRequest{
		TargetId:        "<ENV ID or OU ID>",
		TargetType:      "<Environment/OrganizationalUnit>",
		RulesetId:       "<RuleSetID>",
		NotificationIds: []string{"<NOTIFICATION_IDS>"},
		AdmissionControllerAction:          "<Prevention/Detection>",
        AdmissionControlUnScannedAction:      "<Prevention/Detection>",
	}

	// CreateImageAssurance Policy
	response, _, err := srv.Create(&request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create ImageAssurance Policy response type: %T\n Content: %+v", response, response)

	// Get all ImageAssurance Policy Associations
	allPolicies, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetAll ImageAssurance Policies response type: %T\n Content: %+v", allPolicies, allPolicies)

	// Get specific ImageAssurance Policy
	policy, _, err := srv.Get(response.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get Specific ImageAssurance Policy response type: %T\n Content: %+v", policy, policy)

	// Update specific ImageAssurance Policy
	createRequest.Action = "Detection"
	updatedPolicy, _, err := srv.Update(&createRequest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Update ImageAssurance Policy response type: %T\n Content: %+v", updatedPolicy, updatedPolicy)

	// Delete AC Policy
	deleteResponse, err := srv.Delete(response.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ImageAssurance Policy deleted response type: %T\n Content: %+v", deleteResponse, deleteResponse)
}


```