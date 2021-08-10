```go
package main

import (
    "fmt"
    "github.com/dome9/dome9-sdk-go/dome9"
    "github.com/dome9/dome9-sdk-go/services/serviceaccounts"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := serviceaccounts.New(config)
	req := serviceaccounts.ServiceAccountRequest {
		Name:"YOUR NAME",
		RoleIds:[]int64{1234,5678},
	}

	// Create Service Account
	v, _, err := srv.Create(&req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create response type: %T\n Content %+v", v, v)
}
```