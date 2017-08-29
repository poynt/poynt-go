package poynt

import (
	"fmt"
	"net/url"
)

// CatalogQuery struct for
type TerminalidQuery struct {
	Tid        string `json:"tid,omitempty"`
	DeviceId   string `json:"deviceId,omitempty"`
	BusinessId string `json:"businessId,omitempty"`
	StoreId    string `json:"storeId,omitempty"`
}

type Terminalid interface{}

type TerminalidSingleStore struct {
	Id string `json:"id"`
}

type TerminalidSingle struct {
	Store *TerminalidSingleStore `json:"store"`
}

// GetTerminalid returns a single terminalid
func (api *PoyntApi) GetTerminalid(terminalidId string) (*TerminalidSingle, error) {
	path := fmt.Sprintf("/terminal-ids/%s", terminalidId)

	terminalid := new(TerminalidSingle)
	err := api.Get(path, nil, terminalid)

	if err != nil {
		return nil, err
	}

	return terminalid, nil
}

// GetTerminalids returns the catalogs belonging to this business
func (api *PoyntApi) GetTerminalids(query *TerminalidQuery) ([]*Terminalid, error) {
	path := fmt.Sprintf("/terminal-ids")

	terminalids := new([]*Terminalid)

	urlValues := url.Values{}

	if query.BusinessId != "" {
		urlValues.Set("businessId", query.BusinessId)
	}
	if query.StoreId != "" {
		urlValues.Set("storeId", query.StoreId)
	}
	if query.Tid != "" {
		urlValues.Set("tid", query.Tid)
	}
	if query.DeviceId != "" {
		urlValues.Set("deviceId", query.DeviceId)
	}

	err := api.Get(path, urlValues, terminalids)

	if err != nil {
		return nil, err
	}

	return *terminalids, nil
}
