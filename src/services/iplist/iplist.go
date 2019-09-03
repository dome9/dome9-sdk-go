package iplist

import (
	"fmt"
	"net/http"
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

func (ipLists *ServiceInitiator) Get(ipListId int) (*IpList, *http.Response, error) {
	v := new(IpList)
	path := fmt.Sprintf("iplist/%d", ipListId)
	resp, err := ipLists.Client.NewRequestDo("GET", path, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (ipLists *ServiceInitiator) Create(ipList *IpList) (*IpList, *http.Response, error) {
	v := new(IpList)
	resp, err := ipLists.Client.NewRequestDo("POST", "iplist/", ipList, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (ipLists *ServiceInitiator) Update(ipListId string, ipList *IpList) (*http.Response, error) {
	path := fmt.Sprintf("iplist/%s", ipListId)
	// v is nil because updating iplist returns nothing (204)
	resp, err :=ipLists.Client.NewRequestDo("PUT", path, ipList, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (ipLists *ServiceInitiator) Delete(ipListId string) (*http.Response, error) {
	path := fmt.Sprintf("iplist/%s", ipListId)
	// v is nil because deleting iplist returns nothing (204)
	resp, err := ipLists.Client.NewRequestDo("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}
