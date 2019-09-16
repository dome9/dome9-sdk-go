package cloudaccount

import (
	"net/http"
	"time"
)

// AzureCloudAccountRequest and AzureCloudAccountResponse refer to API type: AzureCloudAccount
type AzureCloudAccountRequest struct {
	Name           string `json:"name"`
	SubscriptionID string `json:"subscriptionId"`
	TenantID       string `json:"tenantId"`
	Credentials    struct {
		ClientID       string `json:"clientId"`
		ClientPassword string `json:"clientPassword"`
	} `json:"credentials"`
	OperationMode          string    `json:"operationMode"`
	Error                  *string   `json:"error"`
	CreationDate           time.Time `json:"creationDate"`
	OrganizationalUnitID   string    `json:"organizationalUnitId"`
	OrganizationalUnitPath string    `json:"organizationalUnitPath"`
	OrganizationalUnitName string    `json:"organizationalUnitName"`
	Vendor                 string    `json:"vendor"`
}

type AzureCloudAccountResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SubscriptionID string `json:"subscriptionId"`
	TenantID       string `json:"tenantId"`
	Credentials    struct {
		ClientID       string  `json:"clientId"`
		ClientPassword *string `json:"clientPassword"`
	} `json:"credentials"`
	OperationMode          string    `json:"operationMode"`
	Error                  *string   `json:"error"`
	CreationDate           time.Time `json:"creationDate"`
	OrganizationalUnitID   *string   `json:"organizationalUnitId"`
	OrganizationalUnitPath string    `json:"organizationalUnitPath"`
	OrganizationalUnitName string    `json:"organizationalUnitName"`
	Vendor                 string    `json:"vendor"`
}

func (service *Service) GetCloudAccountAzure(options interface{}) (*AzureCloudAccountResponse, *http.Response, error) {
	v := new(AzureCloudAccountResponse)
	resp, err := service.client.NewRequestDo("GET", D9AzureResourceName, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) CreateCloudAccountAzure(body interface{}) (*AzureCloudAccountResponse, *http.Response, error) {
	v := new(AzureCloudAccountResponse)
	resp, err := service.client.NewRequestDo("POST", D9AzureResourceName, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
