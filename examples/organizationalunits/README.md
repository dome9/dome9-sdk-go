```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/organizationalunits"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := organizationalunits.New(config)

	request := organizationalunits.OURequest{
		Name:     "tester",
		ParentID: "12345678-1234-1234-1234-12345678901",
	}

    // create OU
	createdOU, _, err := srv.Create(&request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create response type: %T\n Content: %+v", createdOU, createdOU)
    
    // get all OUs
	allOUs, _, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("GetAll response type: %T\n Content: %+v", allOUs, allOUs)
	
    // Get specific OU
	someIpList, _, err := srv.Get("12345678-1234-1234-1234-12345678901")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", someIpList, someIpList)
    
    // update specific OU
	v, err := srv.Update("12345678-1234-1234-1234-12345678901", &request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update response type: %T\n Content: %+v", v, v)

    // Delete IP List
    _, err = srv.Delete("12345678-1234-1234-1234-12345678901")
    if err != nil {
        panic(err)
    }

    fmt.Printf("OU deleted")
}

```