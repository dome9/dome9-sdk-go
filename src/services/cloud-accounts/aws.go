package cloud_accounts

import (
	"dome9"
	"dome9/client"
	"fmt"
	"net/http"
	"time"
)

type Credentials struct {
	ApiKey     string `json:"apikey"`
	Arn        string `json:"arn"`
	Secret     string `json:"secret"`
	IamUser    string `json:"iamUser"`
	Type       string `json:"type"`
	IsReadOnly bool   `json:"isReadOnly"`
}

type Region struct {
	Region           string `json:"region"`
	Name             string `json:"name"`
	Hidden           bool   `json:"hidden"`
	NewGroupBehavior string `json:"newGroupBehavior"`
}

type NetSec struct {
	Regions []Region `json:"regions"`
}

type AWS struct {
	ID                     string      `json:"id"`
	Vendor                 string      `json:"vendor"`
	Name                   string      `json:"name"`
	ExternalAccountNumber  string      `json:"externalAccountNumber"`
	Error                  string      `json:"error"`
	IsFetchingSuspended    bool        `json:"isFetchingSuspended"`
	CreationDate           time.Time   `json:"creationDate"`
	Credentials            Credentials `json:"credentials"`
	IamSafe                string      `json:"iamSafe"`
	NetSec                 NetSec      `json:"netSec"`
	Magellan               bool        `json:"magellan"`
	FullProtection         bool        `json:"fullProtection"`
	AllowReadOnly          bool        `json:"allowReadOnly"`
	OrganizationalUnitID   string      `json:"organizationalUnitId"`
	OrganizationalUnitPath string      `json:"organizationalUnitPath"`
	OrganizationalUnitName string      `json:"organizationalUnitName"`
	LambdaScanner          bool        `json:"lambdaScanner"`
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

// Required properties for onBoarding process
type OnBoarding struct {
	Name              string `json:"name"`
	CustomCredentials struct {
		Arn    string `json:"arn"`
		Secret string `json:"secret"`
		Type   string `json:"type"`
	} `json:"credentials"`
	FullProtection bool `json:"fullProtection"`
	AllowReadOnly  bool `json:"allowReadOnly"`
}

func (aws *Service) Create(onBoarding *OnBoarding) (*AWS, *http.Response, error) {
	v := new(AWS)
	resp, err := aws.client.NewRequestDo("POST", "CloudAccounts/", onBoarding, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (aws *Service) Get(accountId string) (*AWS, *http.Response, error) {
	v := new(AWS)
	path := fmt.Sprintf("cloudaccounts/%s", accountId)
	resp, err := aws.client.NewRequestDo("GET", path, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (aws *Service) Delete(accountId string) (*http.Response, error) {
	path := fmt.Sprintf("cloudaccounts/%s", accountId)
	resp, err := aws.client.NewRequestDo("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
