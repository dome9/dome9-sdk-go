```go
package main

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/awp/azure_onboarding"
	"github.com/dome9/dome9-sdk-go/services/awp"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("Access-ID", "Secret-Key", "https://api.dome9.com/v2/")
	srv := awp_azure_onboarding.New(config)

	cloudAccountId := "<Cloudguard acconut id>"
	
	// Define the request
	awpAzureOnboardingDataRequest := awp_azure_onboarding.CreateAWPOnboardingDataRequest{
		CentralizedId:     "string", // optional 
	}

	// Get awp azure onboarding data
	appAzureOnboardingDataResponse, _, err := srv.GetOnboardingData(cloudAccountId, awpAzureOnboardingDataRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get AWP Azure Onboarding Data response type: %T\n Content: %+v", appAzureOnboardingDataResponse, appAzureOnboardingDataResponse)


	// Define the request
	awpAzureOnboardingRequest := awp_azure_onboarding.CreateAWPOnboardingRequestAzure{
		ScanMode:                   "ScanMode", // can be "inAccount", inAccountHub, inAccountSub or "saas" 
		IsTerraform:                true,
		CentralizedCloudAccountId:  "string", // relevant for inAccountSub mode
		ManagementGroupId:          "string", // relevant for inAccountHub mode
		AgentlessAccountSettings: &awp_onboarding.AgentlessAccountSettings{
			DisabledRegions:              []string{"eastus"},
			ScanMachineIntervalInHours:   26,
			MaxConcurrenceScansPerRegion: 20,
			SkipFunctionAppsScan:         false,
			CustomTags:                   map[string]string{"test": "testValue"},
		},
	}

	// Create AWP Azure Onboarding

	options := awp_onboarding.CreateOptions{
		ShouldCreatePolicy: "true",
	}

	awpAzureOnboardingResponse, err := srv.CreateAWPOnboarding(cloudAccountId, awpAzureOnboardingRequest, options)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Onboarding AWP Azure Account with for type: %T\n Content: %+v", awpAzureOnboardingResponse, awpAzureOnboardingResponse)

	// Get AWP Azure Onboarding

	getAwpOnboardingDataResponse, _, err := srv.GetAWPOnboarding(cloudAccountId)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Get AWP Onbording account configuration response type: %T\n Content: %+v", getAwpOnboardingDataResponse, getAwpOnboardingDataResponse)

	// Define the new settings
	newSettings := awp_onboarding.AgentlessAccountSettings{
		DisabledRegions:              []string{"eastus", "westus"},
		ScanMachineIntervalInHours:   10,
		MaxConcurrenceScansPerRegion: 6,
		SkipFunctionAppsScan:         false,
		CustomTags:                   map[string]string{"newTag": "newValue"},
	}

	// Update AWP Azure Onboarding settings
	updateResponse, err := srv.UpdateAWPSettings(cloudAccountId, newSettings)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Update AWP Azure Onboarding settings response type: %T\n Content: %+v", updateResponse, updateResponse)

	// Get the updated AWP Onboarding data
	updatedAwpOnboardingDataResponse, _, err := srv.GetAWPOnboarding(cloudAccountId)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Get updated AWP Onboarding account configuration response type: %T\n Content: %+v", updatedAwpOnboardingDataResponse, *updatedAwpOnboardingDataResponse.AgentlessAccountSettings)

	
	// Delete AWP Azure Onboarding
	deleteResponse, err := srv.DeleteAWPOnboarding(cloudAccountId)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Delete AWP Azure Onboarding response type: %T\n Content: %+v", deleteResponse, deleteResponse)

}
```