```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupaws"
)

func main() {
	config, _ := dome9.NewConfig("", "", "")
	srv := securitygroupaws.New(config)

	scope := securitygroupaws.Scope{
		Type: "CIDR",
		Data: map[string]string{
			"cidr": "0.0.0.0/0",
			"note": "Allow All Traffic",
		},
	}

	inbound := securitygroupaws.BoundServicesRequest{
		Name:         "NAME",
		Description:  "DESCRIPTION",
		ProtocolType: "TCP",
		Port:         "12",
		OpenForAll:   true,
		Scope: []securitygroupaws.Scope{
			scope,
		},
	}

	outbound := securitygroupaws.BoundServicesRequest{
		Name:         "Name",
		Description:  "DESCRIPTION",
		ProtocolType: "ALL",
		Port:         "",
		OpenForAll:   true,
	}

	req := securitygroupaws.CloudSecurityGroupRequest{
		IsProtected:       true,
		SecurityGroupName: "Name",
		Description:       "DESCRIPTION",
		RegionId:          "us_west_1",
		CloudAccountId:    "00000000-0000-0000-0000-000000000000",
		Services: securitygroupaws.ServicesRequest{
			Inbound: []securitygroupaws.BoundServicesRequest{
				inbound,
			},
			Outbound: []securitygroupaws.BoundServicesRequest{
				outbound,
			},
		},
		Tags: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// create AWS SG
	v, _, err := srv.Create(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create response type: %T\n Content %+v", v, v)

	// get all AWS SG
	allAWSSecurityGroups, _, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("GetAll response type: %T\n Content: %+v", allAWSSecurityGroups, allAWSSecurityGroups)

	// get a specific AWS SG
	someAWSSecurityGroup, _, err := srv.Get("00000000-0000-0000-0000-000000000000", "us_west_1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v", someAWSSecurityGroup, someAWSSecurityGroup)

	securityGroupResponse, _, err := srv.Update("sg-00000000000000000", securitygroupaws.CloudSecurityGroupRequest{
		IsProtected:       true,
		SecurityGroupName: "NAME",
		Description:       "DESCRIPTION",
		RegionId:          "us_west_1",
		CloudAccountId:    "00000000-0000-0000-0000-000000000000",
		Services: securitygroupaws.ServicesRequest{
			Inbound: []securitygroupaws.BoundServicesRequest{
			},
			Outbound: []securitygroupaws.BoundServicesRequest{
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Update AWS security group: %T\n Content: %+v\n", securityGroupResponse, securityGroupResponse)

	// update protection mode
	securityGroupResponse, _, err = srv.UpdateProtectionMode("sg-00000000000000000", "FullManage")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Update protection mode mode: %T\n Content: %+v\n", securityGroupResponse, securityGroupResponse)

	// update bound services
	securityGroupResponse, _, err = srv.HandelBoundServices("sg-00000000000000000", "Inbound", inbound)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Update bound services: %T\n Content: %+v\n", securityGroupResponse, securityGroupResponse)

	// delete AWS SG
	_, err = srv.Delete("sg-00000000000000000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("AWS Security Group deleted")
}

```