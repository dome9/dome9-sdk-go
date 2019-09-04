package cloud_accounts

import (
	"dome9"
	"dome9/client"
	"net/http"
)

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) Create(resourceName string, onBoarding, v interface{}) (interface{}, *http.Response, error) {
	resp, err := service.client.NewRequestDo("POST", resourceName, nil, onBoarding, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(resourceName string, options, v interface{}) (interface{}, *http.Response, error) {
	resp, err := service.client.NewRequestDo("GET", resourceName, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(resourceName string, options interface{}) (*http.Response, error) {
	resp, err := service.client.NewRequestDo("DELETE", resourceName, options,nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
