# dome9-sdk-go
NOTE! This repositoy is currently private. Once it will be public, installing instructions and imports will be changed.
## Aim and Scope
Dome9 GO SDK aims to access Dome9 Web API through HTTP calls
from a client application purely written in Go language.

For more information about Dome9 Web API [Dome9 API](https://api-v2-docs.dome9.com/).

## Pre-requisites
* The API is built using Go 1.12.9. Some features may not be
available or supported unless you have installed a relevant version of Go.
Please click [https://golang.org/dl/](https://golang.org/dl/) to download and
get more information about installing Go on your computer.
* Make sure you have properly set both `GOROOT` and `GOPATH`
environment variables.

* Before you begin, make sure you have an account in [Dome9](https://secure.dome9.com/).

* When logged in, go to My Settings and create your API KEY. Copy the `ID` and `Secret`, they will use you for authetication.

## Installation
To download all packages in the repo with their dependencies, simply run

`go get github.com/Dome9/Dome9-sdk-go/...`

## Getting Started
One can start using Dome9 Go SDK by initializing client and making a request. 
Here is an example of getting IP List.
```go
package main

import (
	"fmt"
	"dome9"
	"services/iplist"
)


func main() {
	accessID := "123456789123456789" // your API ID
	secretKey := "passpasspass"      // you API Secret
	config := dome9.DefaultConfig()
	config.SetKeys(accessID, secretKey)

	srv := iplist.New(config)

	res, _, _ := srv.Get(77281)
	fmt.Println(res)


}
```