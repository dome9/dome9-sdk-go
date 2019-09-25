package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/iplist"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := iplist.New(config)
	var req iplist.IpList

	// Set IP List parameters
	req.Name = "test IP list"
	req.Description = "test description"
	item := struct {
		Ip      string
		Comment string
	}{
		Ip:      "6.6.6.6",
		Comment: "This is comment",
	}
	req.Items = append(req.Items, item)

	v, _, err := srv.Create(&req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Create response type: %T\n Content %+v", v, v)

	resp, _, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v", resp, resp)
}
