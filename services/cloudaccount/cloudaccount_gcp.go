package cloudaccount

import (
	"net/http"
	"time"
)

// refer to API type: GoogleCloudAccountPost
type GCPCloudAccountRequest struct {
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
type GCPCloudAccountResponse struct {
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

func (service *Service) GetCloudAccountGCP(options interface{}) (*GCPCloudAccountResponse, *http.Response, error) {
	v := new(GCPCloudAccountResponse)
	resp, err := service.client.NewRequestDo("GET", D9GCPResourceName, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) CreateCloudAccountGCP(body interface{}) (*GCPCloudAccountResponse, *http.Response, error) {
	v := new(GCPCloudAccountResponse)
	resp, err := service.client.NewRequestDo("POST", D9GCPResourceName, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
