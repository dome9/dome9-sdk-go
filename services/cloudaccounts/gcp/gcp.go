package gcp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Dome9/dome9-sdk-go/services/cloudaccounts"
)

// refer to API type: GoogleCloudAccountPost
type CloudAccountRequest struct {
	Name                      string `json:"name"`
	ServiceAccountCredentials struct {
		Type                    string `json:"type"`
		ProjectID               string `json:"project_id"`
		PrivateKeyID            string `json:"private_key_id"`
		PrivateKey              string `json:"private_key"`
		ClientEmail             string `json:"client_email"`
		ClientID                string `json:"client_id"`
		AuthURI                 string `json:"auth_uri"`
		TokenURI                string `json:"token_uri"`
		AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
		ClientX509CertURL       string `json:"client_x509_cert_url"`
	} `json:"serviceAccountCredentials"`
	GsuiteUser string `json:"gsuiteUser"`
	DomainName string `json:"domainName"`
}

// refer to API type: GoogleCloudAccountGet
type CloudAccountResponse struct {
	ID                     string    `json:"id"`
	Name                   string    `json:"name"`
	ProjectID              string    `json:"projectId"`
	CreationDate           time.Time `json:"creationDate"`
	OrganizationalUnitID   *string   `json:"organizationalUnitId"`
	OrganizationalUnitPath string    `json:"organizationalUnitPath"`
	OrganizationalUnitName string    `json:"organizationalUnitName"`
	Gsuite                 *struct {
		GsuiteUser string `json:"gsuiteUser"`
		DomainName string `json:"domainName"`
	} `json:"gsuite"`
	Vendor string `json:"vendor"`
}

func (service *Service) Get(options interface{}) (*CloudAccountResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathGCP, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathGCP, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathGCP, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeAddress := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathGCP, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeAddress, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
