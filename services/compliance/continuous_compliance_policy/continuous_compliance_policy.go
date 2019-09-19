package continuous_compliance_policy

import (
	"fmt"
	"net/http"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/dome9/client"
)

const path = "Compliance/ContinuousCompliancePolicy/"

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

type ContinuousCompliancePolicyRequest struct {
	CloudAccountID    string   `json:"cloudAccountId"`
	ExternalAccountID string   `json:"externalAccountId"`
	CloudAccountType  string   `json:"cloudAccountType"`
	BundleID          int      `json:"bundleId"`
	NotificationIds   []string `json:"notificationIds"`
}

type ContinuousCompliancePolicyResponse struct {
	ID                string   `json:"id"`
	CloudAccountID    string   `json:"cloudAccountId"`
	ExternalAccountID string   `json:"externalAccountId"`
	CloudAccountType  string   `json:"cloudAccountType"`
	BundleID          int      `json:"bundleId"`
	NotificationIds   []string `json:"notificationIds"`
}

func (service *Service) Get(options interface{}) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}
	v := new(ContinuousCompliancePolicyResponse)
	resp, err := service.client.NewRequestDo("GET", path, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new([]ContinuousCompliancePolicyResponse)
	resp, err := service.client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *ContinuousCompliancePolicyRequest) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new(ContinuousCompliancePolicyResponse)
	resp, err := service.client.NewRequestDo("POST", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(id string, body *ContinuousCompliancePolicyRequest) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new(ContinuousCompliancePolicyResponse)
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
