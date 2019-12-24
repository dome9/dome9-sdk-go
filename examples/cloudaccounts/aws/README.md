```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
)

func main() {
	// pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := aws.New(config)
	var req aws.CloudAccountRequest

	reqIamSafe := aws.AttachIamSafeRequest{
		CloudAccountID: "000000-0000-0000-0000-00000000000000",
		Data: aws.Data{
			AwsGroupArn:  "GROUP-ARN",
			AwsPolicyArn: "POLICY-ARN",
		},
	}

	// must fill below variables
	req.Name = "test AWS cloud account"
	req.Credentials.Type = "RoleBased"
	req.Credentials.Arn = "ARN"
	req.Credentials.Secret = "SECRET"

	// create cloud account
	v, _, err := srv.Create(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create response type: %T\n Content %+v\n", v, v)

	// get all cloud accounts
	cloudAccounts, _, err := srv.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v\n", cloudAccounts, cloudAccounts)

	// get specific cloud accounts
	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: "SOME_ID"}
	cloudAccount, _, err := srv.Get(&getCloudAccountQueryParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get response type: %T\n Content: %+v\n", cloudAccount, cloudAccount)

	// update cloud account name
	desiredNewName := "test AWS update cloud account"
	updateNameResponse, _, err := srv.UpdateName(aws.CloudAccountUpdateNameRequest{
		CloudAccountID: v.ID,
		Data:           desiredNewName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("update name response type: %T\n Content: %+v\n", updateNameResponse, updateNameResponse)

	// update region config
	desiredGroupBehavior := "FullManage"
	updateRegionConfigResponse, _, err := srv.UpdateRegionConfig(aws.CloudAccountUpdateRegionConfigRequest{
		CloudAccountID: v.ID,
		Data: aws.CloudAccountNetSecRegion{
			Region:           "us_east_1",
			NewGroupBehavior: desiredGroupBehavior,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("update region config response type: %T\n Content: %+v\n", updateRegionConfigResponse, updateRegionConfigResponse)

	// update organizational unit id
	OrganizationalUnitID := "ORGANIZATIONAL_UNIT_ID"
	id := "THE ACCOUNT ID IN DOME9"
	updateOrganizationalIDResponse, _, err := srv.UpdateOrganizationalID(id,
		aws.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitId: OrganizationalUnitID},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("update Organizational ID response type: %T\n Content: %+v\n", updateOrganizationalIDResponse, updateOrganizationalIDResponse)

	// update credentials
	updateCredentialsResponse, _, err := srv.UpdateCredentials(aws.CloudAccountUpdateCredentialsRequest{
		CloudAccountID: v.ID,
		Data: aws.CloudAccountCredentials{
			Arn:    "ARN",
			Secret: "SECRET",
			Type:   "RoleBased",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Update credentials response type: %T\n Content: %+v\n", updateCredentialsResponse, updateCredentialsResponse)

	// delete aws cloud account
	_, err = srv.Delete("SOME_ID")
	if err != nil {
		panic(err)
	}
	fmt.Printf("AWS cloud accout deleted")

	// attach iam safe to cloud account
	v, _, err = srv.AttachIAMSafeToCloudAccount(reqIamSafe)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Attach response type: %T\n Content %+v", *v, *v)

	// detach iam safe
	_, err = srv.DetachIAMSafeToCloudAccount("000000-0000-0000-0000-00000000000000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Detach IAM safe")

	var attachRestrictedIamEntitiesResp *string
	restrictedIamEntitiesRequest := aws.RestrictedIamEntitiesRequest{
		EntityName: "AWS_IAM_USER_NAME",
		EntityType: "User",
	}
	
	attachRestrictedIamEntitiesResp , _, err = srv.ProtectAWSIAMEntity("000000-0000-0000-0000-00000000000000", restrictedIamEntitiesRequest)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Printf("Attach restricted IAM entities. Role arn %s", *attachRestrictedIamEntitiesResp)
    
    _, err = srv.UnprotectAWSIAMEntity("000000-0000-0000-0000-00000000000000", "AWS_IAM_USER_NAME", "User")
	if err != nil {
		fmt.Println(err)
	}
    fmt.Printf("Detach restricted IAM entities")

    restrictedUserIamEntities, err := srv.GetProtectAWSIAMEntityStatusByName("000000-0000-0000-0000-00000000000000", "USER_NAME", "User")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Restricted user iam entities: %T\n Content: %+v\n", *restrictedUserIamEntities, *restrictedUserIamEntities)

}

```