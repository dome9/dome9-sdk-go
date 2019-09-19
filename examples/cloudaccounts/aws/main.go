package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts/aws"
)

func main() {
	// Pass accessID, secretKey or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := aws.New(config)
	var req aws.CloudAccountRequest

	req.Name = "test AWS cloud account"
	req.FullProtection = false
	req.AllowReadOnly = false
	req.Vendor = "aws"
	arn := "ARN"
	secret := "SECRET"
	req.Credentials.Arn = &arn
	req.Credentials.Secret = secret
	req.Credentials.Type = "RoleBased"

	v, _, err := srv.Create(req)
	resp, _, _ := srv.GetAll()
	fmt.Printf("Create response type: %T\n Content %+v\n", v, v)
	fmt.Printf("Get response type: %T\n Content: %+v\n", resp, resp)

	if err != nil {
		fmt.Println(err)
	}

}
