package continuous_compliance_policy

import (
	"net/http"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/dome9/client"
)

const path = "Compliance/ContinuousCompliancePolicy/"

type ContinuousCompliancePolicyGet struct {
	ID                string   `json:"id"`
	CloudAccountID    string   `json:"cloudAccountId"`
	ExternalAccountID string   `json:"externalAccountId"`
	CloudAccountType  string   `json:"cloudAccountType"`
	BundleID          int      `json:"bundleId"`
	NotificationIds   []string `json:"notificationIds"`
}

type ContinuousCompliancePolicyPut struct {
	CloudAccountID    string   `json:"cloudAccountId"`
	ExternalAccountID string   `json:"externalAccountId"`
	CloudAccountType  string   `json:"cloudAccountType"`
	BundleID          int      `json:"bundleId"`
	NotificationIds   []string `json:"notificationIds"`
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) Create(body *ContinuousCompliancePolicyPut) (*ContinuousCompliancePolicyGet, *http.Response, error) {
	v := new(ContinuousCompliancePolicyGet)
	resp, err := service.client.NewRequestDo("POST", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() ([]ContinuousCompliancePolicyGet, *http.Response, error) {
	var v []ContinuousCompliancePolicyGet
	resp, err := service.client.NewRequestDo("GET", path, nil, nil, &v)

	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(id string, body *ContinuousCompliancePolicyPut) (*ContinuousCompliancePolicyGet, *http.Response, error) {
	v := new(ContinuousCompliancePolicyGet)
	resp, err := service.client.NewRequestDo("PUT", path+id, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	resp, err := service.client.NewRequestDo("DELETE", path+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
