package main

import (
	"fmt"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/services/assessment"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := assessment.New(config)
	var req assessment.RunBundleRequest

	req.ID = 89283
	req.Dome9CloudAccountID = "DOME_9_CLOUD_ACCOUNT_ID"

	v, _, err := srv.RunBundle(&req)
	fmt.Printf("Run bundle response type: %T\n Content %+v", v, v)

	if err != nil {
		fmt.Println(err)
	}

}
