```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/iplist"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := iplist.New(config)

	request := iplist.IpList{
		Name:        "test IP list",
		Description: "test description",
		Items: []iplist.Item{
			{
				Ip:      "6.6.6.6",
				Comment: "This is comment",
			},
		},
	}

    // Create IP List
	createdIpList, _, err := srv.Create(&request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create response type: %T\n Content: %+v", createdIpList, createdIpList)
    
    // Get all IP Lists
	allIpLists, _, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("GetAll response type: %T\n Content: %+v", allIpLists, allIpLists)
	
    // Get specific IP List
	someIpList, _, err := srv.Get(10001)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", someIpList, someIpList)
    
    // Update specific IP List
	v, err := srv.Update(10001, &request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update response type: %T\n Content: %+v", v, v)

    // Delete IP List
    _, err := srv.Delete(1001)
    if err != nil {
        panic(err)
    }

    fmt.Printf("IP List deleted")
}

```