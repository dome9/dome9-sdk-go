# dome9-sdk-go

## Aim and Scope
Dome9 GO SDK aims to access Dome9 Web API through HTTPS calls
from a client application purely written in Go language.

For more information about Dome9 Web API [Dome9 API](https://api-v2-docs.dome9.com/).

## Prerequisites
* The API is built using Go 1.19. Some features may not be
available or supported unless you have installed a relevant version of Go.
Please click [https://golang.org/dl/](https://golang.org/dl/) to download and
get more information about installing Go on your computer.
* Make sure you have properly set both `GOROOT` and `GOPATH`
environment variables.

* Before you begin, make sure you have an account in [Dome9](https://secure.dome9.com/).

* When logged in, go to My Settings and create your API KEY. Copy the `accessID` and `secretKey`, they will use you for authentication.

## Installation
To download all packages in the repo with their dependencies, simply run

`go get github.com/dome9/dome9-sdk-go`

## Getting Started
One can start using Dome9 Go SDK by initializing client and making a request. 
Here is an example of getting IP List.

```go
package main

import (
	"fmt"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/iplist"
)

func main() {
	accessID := "ACCESS ID"
	secretKey := "SECRET KEY"
	config, err := dome9.NewConfig(accessID, secretKey, "")

	if err != nil {
		panic(err)
	}

	ipListService := iplist.New(config)
	var ipListID int64 = 77281
	response, _, err := ipListService.Get(ipListID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Received ipLists: %+v", response)
}
```
