package main


import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"

	"github.com/dome9/dome9-sdk-go/dome9"
)

func main() {
	accessID := "ada7d502-bd76-43fe-adee-4ef14c15df07"
	secretKey := "u1oqhmbqb8kddugg1hofuwjz"
	config, err := dome9.NewConfig(accessID, secretKey, "")

	srv := continuous_compliance_policy.New(config)
	var req continuous_compliance_policy.ContinuousCompliancePolicyRequest
	req.CloudAccountType = "Aws"

	// Set Rule bundle ID
	desiredBundleID := 86685
	req.BundleID = desiredBundleID

	// must fill below variables
	req.CloudAccountID = "98f001ae-abd2-4f03-9d67-ab80b84498f1"
	req.ExternalAccountID = "QLe7HCZX@zQD+og@zr3ALKCM"
	req.NotificationIds = []string{"df1b6f0f-966c-4cbf-91bc-5b7b0ebd306b"}
/*
	// Create CC Policy
	v, _, err := srv.Create(&req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create response type: %T\n Content %+v", v, v)
*/
	// Get all CC Policies
	resp, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", resp, resp)
/*
	// Get specific CC Policy
	somePolicy, _, err := srv.Get("877cbf1f-8f03-43d6-a1b9-35b61252fc9a")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", somePolicy, somePolicy)

	// Update specific CC Policy
	v, _, err = srv.Update("877cbf1f-8f03-43d6-a1b9-35b61252fc9a", &req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update response type: %T\n Content: %+v", v, v)

	// Delete CC Policy
	_, err = srv.Delete("877cbf1f-8f03-43d6-a1b9-35b61252fc9a")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Continuous Compliance Policy deleted")

 */
}