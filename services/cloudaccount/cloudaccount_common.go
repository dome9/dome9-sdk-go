package cloudaccount

import (
	"net/http"

	"github.com/Dome9/dome9-sdk-go/dome9"
	"github.com/Dome9/dome9-sdk-go/dome9/client"
)

const (
	D9AwsResourceName   = "cloudaccounts/"
	D9AzureResourceName = "AzureCloudAccount/"
	D9GCPResourceName   = "GoogleCloudAccount/"
)

type QueryParameters struct {
	ID string
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{client: client.NewClient(c)}
}

func (service *Service) Delete(resourceNamePath string, id string) (*http.Response, error) {
	resp, err := service.client.NewRequestDo("DELETE", resourceNamePath+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
