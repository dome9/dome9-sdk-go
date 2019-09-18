package azure

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts"
)

// AzureCloudAccountRequest and CloudAccountsResponse refer to API type: AzureCloudAccount
type CloudAccountsRequest struct {
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

type CloudAccountsResponse struct {
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

func (service *Service) Get(options interface{}) (*CloudAccountsResponse, *http.Response, error) {
	v := new(CloudAccountsResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAzure, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountsRequest) (*CloudAccountsResponse, *http.Response, error) {
	v := new(CloudAccountsResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathAzure, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeAddress := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAzure, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeAddress, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
