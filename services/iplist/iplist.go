package iplist

import (
	"fmt"
	"github.com/dome9-sdk-go/dome9"
	"github.com/dome9-sdk-go/dome9/client"
	"net/http"
)

const (
	ipListResourcePath = "iplist"
)

type Ip struct {
	Ip      string
	Comment string
}

type IpList struct {
	Id          int64
	Name        string
	Description string
	Items       []Ip
}

type Service struct {
	client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{
		client: client.NewClient(c),
	}
}

func (ipLists *Service) Get(ipListId int64) (*IpList, *http.Response, error) {
	v := new(IpList)
	path := fmt.Sprintf("%s/%d", ipListResourcePath, ipListId)
	resp, err := ipLists.client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (ipLists *Service) GetAll() (*[]IpList, *http.Response, error) {
	v := new([]IpList)
	resp, err := ipLists.client.NewRequestDo("GET", ipListResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (ipLists *Service) Create(ipList *IpList) (*IpList, *http.Response, error) {
	v := new(IpList)
	resp, err := ipLists.client.NewRequestDo("POST", ipListResourcePath, nil, ipList, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (ipLists *Service) Update(ipListId int64, ipList *IpList) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d", ipListResourcePath, ipListId)
	// v is nil because updating iplist returns nothing (204)
	resp, err := ipLists.client.NewRequestDo("PUT", path, ipList, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (ipLists *Service) Delete(ipListId int64) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d", ipListResourcePath, ipListId)
	// v is nil because deleting iplist returns nothing (204)
	resp, err := ipLists.client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
