```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/accesslease"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "", nil)
	srv := accesslease.New(config)

	var resp *accesslease.Response
	req := accesslease.Request{
		CloudAccountID:  "00000000-0000-0000-0000-000000000000",
		Region:          "us_west_2",
		SecurityGroupID: "0000000",
		Protocol:        "ALL",
	}

	_, err := srv.Create(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("access lease created successfully")

	resp, _, err = srv.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("access lese response type: %T\n Content %+v", *resp, *resp)

	// delete aws cloud account
	_, err = srv.Delete("dd1ad89b-b462-4081-839d-598386250f87")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("access lease deleted successfully")
}

```