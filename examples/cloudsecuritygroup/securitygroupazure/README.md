```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupazure"
)

func main() {
	config, _ := dome9.NewConfig("", "", "", nil)

	srv := securitygroupazure.New(config)

	someTags := securitygroupazure.Tags{
		Key:   "yada",
		Value: "gogo",
	}

	data1 := map[string]string{
		"cidr": "0.0.0.0/0",
		"note": "Any",
	}

	sampleScope := securitygroupazure.Scope{
		Type: "CIDR",
		Data: data1,
	}

	inboundServices := securitygroupazure.BoundService{
		Direction:             "Inbound",
		Name:                  "Port_8080_test",
		Description:           "",
		Priority:              100,
		Access:                "Allow",
		Protocol:              "ANY",
		SourcePortRanges:      []string{"3389"},
		SourceScopes:          []securitygroupazure.Scope{sampleScope},
		DestinationPortRanges: []string{"3389"},
		DestinationScopes:     []securitygroupazure.Scope{sampleScope},
		IsDefault:             false,
	}

	outboundServices := securitygroupazure.BoundService{
		Direction:             "Outbound",
		Name:                  "yada",
		Description:           "bada",
		Priority:              100,
		Access:                "Allow",
		Protocol:              "ANY",
		SourcePortRanges:      []string{"3389"},
		SourceScopes:          []securitygroupazure.Scope{sampleScope},
		DestinationPortRanges: []string{"3389"},
		DestinationScopes:     []securitygroupazure.Scope{sampleScope},
		IsDefault:             false,
	}

	req := securitygroupazure.CloudSecurityGroupRequest{
		Name:              "momo",
		Description:       "benzi",
		Region:            "Central US",
		IsTamperProtected: false,
		ResourceGroup:     "erez-rg",
		CloudAccountID:    "CLOUD_ACCOUNT_ID",
		Tags:              []securitygroupazure.Tags{someTags},
		InboundServices:   []securitygroupazure.BoundService{inboundServices},
		OutboundServices:  []securitygroupazure.BoundService{outboundServices},
	}

	// create Azure SG
	resp, _, err := srv.Create(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create response type: %T\n content: %+v'n", resp, resp)

	// get all Azure SGs
	allAzureSecurityGroups, _, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("GetAll response type: %T\n Content: %+v", allAzureSecurityGroups, allAzureSecurityGroups)

	// get a specific Azure SG
	someAzureSecurityGroup, _, err := srv.Get("CLOUD_ACCOUNT_ID")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", someAzureSecurityGroup, someAzureSecurityGroup)

	// update specific Azure SG
	v, _, err := srv.Update("CLOUD_ACCOUNT_ID", req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Update response type: %T\n Content: %+v", v, v)

	// delete Azure SG
	_, err = srv.Delete("CLOUD_ACCOUNT_ID")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Azure Security Group deleted")
}

```