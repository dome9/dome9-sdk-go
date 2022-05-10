```go
package main

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/admissioncontrol/admission_policy"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("Access-ID", "Secret-Key", "https://api.dome9.com/v2/")
	srv := admission_policy.New(config)

	admissionControlCreateRequest := admission_policy.AdmissionControlPolicyRequest{
		TargetId:        "<K8S_Cluster_ID>",
		TargetType:      "Environment",
		RulesetId:       "<RuleSetID>",
		NotificationIds: []string{"<NOTIFICATION_IDS>"},
		Action:          "Prevention",
		RulesetPlatform: "kubernetesruntimeassurance",
	}

	// CreateAdmission Control Policy
	admissionControlPolicyResponse, _, err := srv.Create(&admissionControlCreateRequest)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create AC Policy response type: %T\n Content: %+v", admissionControlPolicyResponse, admissionControlPolicyResponse)

	// Get all Admission Control Policy Associations
	allAdmissionControlPolicies, _, err := srv.GetAllAdmissionControlPolicies()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("GetAll Admission Control Policies response type: %T\n Content: %+v", allAdmissionControlPolicies, allAdmissionControlPolicies)

	// Get specific Admission Control Policy
	admissionControlPolicy, _, err := srv.GetAdmissionControlPolicy(admissionControlPolicyResponse.ID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Get Specific Admission Control Policy response type: %T\n Content: %+v", admissionControlPolicy, admissionControlPolicy)

	// Update specific Admission Control Policy
	admissionControlCreateRequest.Action = "Detection"
	admissionControlUpdatedPolicy, _, err := srv.UpdateAdmissionControlPolicy(&admissionControlCreateRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Update AC Policy response type: %T\n Content: %+v", admissionControlUpdatedPolicy, admissionControlUpdatedPolicy)

	// Delete AC Policy
	deleteResponse, err := srv.Delete(admissionControlPolicyResponse.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("AC Policy deleted response type: %T\n Content: %+v", deleteResponse, deleteResponse)
}


```