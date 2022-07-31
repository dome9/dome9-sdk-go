package continuous_compliance_finding

import (
	"fmt"
	"net/http"
	"time"
)

const (
	continuousComplianceResourcePath = "Finding"
)

type Sorting struct {
	FieldName string `json:"fieldName,omitempty"`
	Direction int    `json:"direction,omitempty"`
}

type DateRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

type FieldFilter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Filter struct {
	FreeTextPhrase   string         `json:"freeTextPhrase,omitempty"`
	Fields           *[]FieldFilter `json:"fields,omitempty"`
	OnlyCIEM         bool           `json:"onlyCIEM,omitempty"`
	IncludedFeatures []string       `json:"includedFeatures,omitempty"`
	CreationTime     *DateRange     `json:"creationTime,omitempty"`
}

type ContinuousComplianceFindingRequest struct {
	PageSize     int        `json:"pageSize,omitempty"`
	Sorting      *Sorting   `json:"sorting,omitempty"`
	MultiSorting []*Sorting `json:"multiSorting,omitempty"`
	Filter       *Filter    `json:"filter,omitempty"`
	SearchAfter  []string   `json:"searchAfter,omitempty"`
	DataSource   string     `json:"dataSource,omitempty"`
}

type ContinuousComplianceFindingResponse struct {
}

func (service *Service) Create(body *ContinuousComplianceFindingRequest) (*ContinuousComplianceFindingResponse, *http.Response, error) {
	v := new(ContinuousComplianceFindingResponse)
	path := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
