package iplist

import (
	"dome9"
	"dome9/client"

	"net/http"
)

type Service interface {
	Create(ipList *IpList) (*IpList, *http.Response, error)
	Get(ipListId string) (*IpList, *http.Response, error)
	Update(ipListId string, ipList *IpList) (*http.Response, error)
	Delete(ipListId string) (*http.Response, error)
}


type ServiceInitiator struct {
	Client *client.Client
}

func New(c *dome9.Config) *ServiceInitiator {
	return &ServiceInitiator{
		Client: client.NewClient(c),
	}
}