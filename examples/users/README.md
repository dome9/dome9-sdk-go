```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/users"
)

func main() {
	// Pass accessID, secretKey, rawUrl, or set environment variables
	config, _ := dome9.NewConfig("", "", "")
	srv := users.New(config)

	request := users.UserRequest{
		Email:      "test@tester.kombina",
		FirstName:  "little",
		LastName:   "finger",
		SsoEnabled: false,
	}

    // create user
	createdUser, _, err := srv.Create(&request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Create response type: %T\n Content: %+v", createdUser, createdUser)
    
    // get all users
	allUsers, _, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("GetAll response type: %T\n Content: %+v", allUsers, allUsers)
	
    // get a specific User
	someUser, _, err := srv.Get("10001")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Get response type: %T\n Content: %+v", someUser, someUser)
    
    // update specific user
	v, err := srv.Update(10001, &request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update response type: %T\n Content: %+v", v, v)

    // delete user
    _, err  = srv.Delete("1001")
    if err != nil {
        panic(err)
    }

    fmt.Printf("User deleted")
}

```