package accesslease

import (
	"fmt"
	"net/http"
	"time"
)

const (
	accessLease = "accessLease"
	aws         = "aws"
)

type Request struct {
	CloudAccountID  string    `json:"cloudAccountId"`
	Region          string    `json:"region"`
	SecurityGroupID string    `json:"securityGroupId"`
	Protocol        string    `json:"protocol"`
	Name            string    `json:"name,omitempty"`
	ID              string    `json:"id,omitempty"`
	IP              string    `json:"ip,omitempty"`
	Note            string    `json:"note,omitempty"`
	Created         time.Time `json:"created,omitempty"`
	User            string    `json:"user,omitempty"`
	Length          string    `json:"length,omitempty"`
	PortFrom        int       `json:"portFrom,omitempty"`
	PortTo          int       `json:"portTo,omitempty"`
	Srl             string    `json:"srl,omitempty"`
}

type Response struct {
	Aws    []AWSAccessLeaseResponse `json:"aws"`
	Agents []AgentsAccessResponse   `json:"agents"`
}

type AWSAccessLeaseResponse struct {
	CloudAccountID  string    `json:"cloudAccountId"`
	Region          string    `json:"region"`
	SecurityGroupID int       `json:"securityGroupId"`
	ID              string    `json:"id"`
	AccountID       int       `json:"accountId"`
	Name            string    `json:"name"`
	IP              string    `json:"ip"`
	Note            string    `json:"note"`
	Created         time.Time `json:"created"`
	User            string    `json:"user"`
	Length          string    `json:"length"`
	Protocol        string    `json:"protocol"`
	PortFrom        int       `json:"portFrom"`
	PortTo          int       `json:"portTo"`
	Srl             string    `json:"srl"`
}

type AgentsAccessResponse struct {
	AgentID   int       `json:"agentId"`
	ID        string    `json:"id"`
	AccountID int       `json:"accountId"`
	Name      string    `json:"name"`
	IP        string    `json:"ip"`
	Note      string    `json:"note"`
	Created   time.Time `json:"created"`
	User      string    `json:"user"`
	Length    string    `json:"length"`
	Protocol  string    `json:"protocol"`
	PortFrom  int       `json:"portFrom"`
	PortTo    int       `json:"portTo"`
	Srl       string    `json:"srl"`
}

func (service *Service) Create(body Request) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", accessLease, aws)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) Get() (*Response, *http.Response, error) {
	v := new(Response)
	resp, err := service.Client.NewRequestDo("GET", accessLease, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", accessLease, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
