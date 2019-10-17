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

	createdIpList, _, err := srv.Create(&request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create response type: %T\n Content: %+v", createdIpList, createdIpList)

	allIpLists, _, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("GetAll response type: %T\n Content: %+v", allIpLists, allIpLists)
}
