package azure

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts"
)

// AzureCloudAccountRequest and CloudAccountResponse refer to API type: AzureCloudAccount
type CloudAccountRequest struct {
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

type CloudAccountResponse struct {
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

func (service *Service) Get(options interface{}) (*CloudAccountResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAzure, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAzure, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
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
