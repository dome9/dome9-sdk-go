```go
package main

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/awp_aws_onboarding"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("Access-ID", "Secret-Key", "https://api.dome9.com/v2/")
	srv := awp_aws_onboarding.New(config)

	// Get awp aws onboarding data
	appAwsOnboardingDataResponse, _, err := srv.Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get AWP AWS Onboarding Data response type: %T\n Content: %+v", appAwsOnboardingDataResponse, appAwsOnboardingDataResponse)
}
```