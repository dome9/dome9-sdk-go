package cloud_accounts

import (
	"net/http"
	"time"
)

const awsResourceName = "cloudaccounts"

type awsCredentials struct {
	ApiKey     string `json:"apikey"`
	Arn        string `json:"arn"`
	Secret     string `json:"secret"`
	IamUser    string `json:"iamUser"`
	Type       string `json:"type"`
	IsReadOnly bool   `json:"isReadOnly"`
}

type region struct {
	Region           string `json:"region"`
	Name             string `json:"name"`
	Hidden           bool   `json:"hidden"`
	NewGroupBehavior string `json:"newGroupBehavior"`
}

type netSec struct {
	Regions []region `json:"regions"`
}

type awsProperties struct {
	ID                     string         `json:"id"`
	Vendor                 string         `json:"vendor"`
	Name                   string         `json:"name"`
	ExternalAccountNumber  string         `json:"externalAccountNumber"`
	Error                  string         `json:"error"`
	IsFetchingSuspended    bool           `json:"isFetchingSuspended"`
	CreationDate           time.Time      `json:"creationDate"`
	Credentials            awsCredentials `json:"credentials"`
	IamSafe                string         `json:"iamSafe"`
	NetSec                 netSec         `json:"netSec"`
	Magellan               bool           `json:"magellan"`
	FullProtection         bool           `json:"fullProtection"`
	AllowReadOnly          bool           `json:"allowReadOnly"`
	OrganizationalUnitID   string         `json:"organizationalUnitId"`
	OrganizationalUnitPath string         `json:"organizationalUnitPath"`
	OrganizationalUnitName string         `json:"organizationalUnitName"`
	LambdaScanner          bool           `json:"lambdaScanner"`
}

type awsOptions struct {
	ID string
}

// Required properties for onBoarding process
type AwsOnBoarding struct {
	Name              string `json:"name"`
	CustomCredentials struct {
		Arn    string `json:"arn"`
		Secret string `json:"secret"`
		Type   string `json:"type"`
	} `json:"credentials"`
	FullProtection bool `json:"fullProtection"`
	AllowReadOnly  bool `json:"allowReadOnly"`
}

func (service *Service) AwsCreate(onBoarding *AwsOnBoarding) (interface{}, *http.Response, error) {
	return service.Create(awsResourceName, onBoarding, new(awsProperties))
}

func (service *Service) AwsGet(accountId string) (interface{}, *http.Response, error) {
	o := &awsOptions{ID: accountId}
	return service.Get(awsResourceName, o, new(awsProperties))
}

func (service *Service) AwsDelete(accountId string) (*http.Response, error) {
	o := &awsOptions{ID: accountId}
	return service.Delete(awsResourceName, o)
}
