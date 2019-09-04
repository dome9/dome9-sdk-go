package cloud_accounts

import (
	"net/http"
	"time"
)

const azureResourceName = "AzureCloudAccount"

type AzureCredentials struct {
	ClientID       string `json:"clientId"`
	ClientPassword string `json:"clientPassword"`
}

type azureProperties struct {
	ID                     string           `json:"id"`
	Name                   string           `json:"name"`
	SubscriptionID         string           `json:"subscriptionId"`
	TenantID               string           `json:"tenantId"`
	Credentials            AzureCredentials `json:"credentials"`
	OperationMode          string           `json:"operationMode"`
	Error                  string           `json:"error"`
	CreationDate           time.Time        `json:"creationDate"`
	OrganizationalUnitID   string           `json:"organizationalUnitId"`
	OrganizationalUnitPath string           `json:"organizationalUnitPath"`
	OrganizationalUnitName string           `json:"organizationalUnitName"`
	Vendor                 string           `json:"vendor"`
}

// Required properties for onBoarding process
type AzureOnBoarding struct {
	Name           string           `json:"name"`
	SubscriptionID string           `json:"subscriptionId"`
	TenantID       string           `json:"tenantId"`
	Credentials    AzureCredentials `json:"credentials"`
	OperationMode  string           `json:"operationMode"`
}

type azureOptions struct {
	ID string
}

func (service *Service) AzureCreate(onBoarding *AzureOnBoarding) (interface{}, *http.Response, error) {
	return service.Create(azureResourceName, onBoarding, new(azureProperties))
}

func (service *Service) AzureGet(accountId string) (interface{}, *http.Response, error) {
	o := &azureOptions{ID: accountId}
	return service.Get(azureResourceName, o, new(azureProperties))
}

func (service *Service) AzureDelete(accountId string) (*http.Response, error) {
	o := &azureOptions{ID: accountId}
	return service.Delete(azureResourceName, o)
}
