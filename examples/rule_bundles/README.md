```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/rulebundles"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := rulebundles.New(config)
    
    sampleRule := rulebundles.Rule {
        Name:"",
        Severity:"",
        Logic:"",
        Description:"",
        Remediation:"",
        ComplianceTag:"",
        Domain:"",
        Priority:"",
        ControlTitle:"",
        RuleID:"",
        LogicHash:"",
        IsDefault:false,
    }
    
	request := rulebundles.RuleBundleRequest{
        Name:"someRuleSet",
        Description:"bla bla",
        Rules: &[]rulebundles.Rule {sampleRule},
        ID:0,
        HideInCompliance:false,
        MinFeatureTier:"",
        CloudVendor:"",
        Language:"",
    }

    // Create Rule Bundle
	createdRuleBundle, _, err := srv.Create(&request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create response type: %T\n Content: %+v", createdRuleBundle, createdRuleBundle)
    
    // Get all Rule Bundles associated to an account
	allIpLists, _, err := srv.GetAccountRuleBundles()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("GetAll response type: %T\n Content: %+v", allIpLists, allIpLists)

    // Get specific Rule Bundle
    someRuleBundle, _, err := srv.Get("SOME_ID")
    if err != nil {
        fmt.Println(err)
    }

	fmt.Printf("Get response type: %T\n Content: %+v", someRuleBundle, someRuleBundle)
    
    // Update specific Rule Bundle
	v, _, err := srv.Update(&request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update response type: %T\n Content: %+v", v, v)

    // Delete Rule bundle
    _, err := srv.Delete("SOME_ID")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Rule bundle deleted")
}   

```