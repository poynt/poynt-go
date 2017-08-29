package poynt

import (
	"fmt"
	"time"
)

type Category interface{}
type Product interface{}
type Discounts interface{}
type Taxes interface{}

// CatalogQuery struct for
type CatalogQuery struct {
	Id         string `json:"catalogId,omitempty"`
	BusinessId string `json:"businessId,omitempty"`
	StoreId    string `json:"storeId,omitempty"`
}

type Catalog struct {
	Id                 string       `json:"id"`
	BusinessId         string       `json:"businessId,omitempty"`
	Categories         []*Category  `json:"categories"`
	Products           []*Product   `json:"products"`
	AvailableDiscounts []*Discounts `json:"availableDiscounts"`
	Taxes              []*Taxes     `json:"taxes"`
	Name               string       `json:"name"`
	CreatedAt          time.Time    `json:"createdAt"`
	UpdatedAt          time.Time    `json:"updatedAt"`
}

type Catalogs struct {
	Catalogs []*Catalog
}

// GetCatalog returns a single catalog
func (api *PoyntApi) GetCatalog(businessId string, catalogId string) (*Catalog, error) {
	path := fmt.Sprintf("/businesses/%s/catalogs/%s/full", businessId, catalogId)

	catalog := new(Catalog)
	err := api.Get(path, nil, catalog)

	if err != nil {
		return nil, err
	}

	return catalog, nil
}

// GetCatalogs returns the catalogs belonging to this business
func (api *PoyntApi) GetCatalogs(businessId string) ([]*Catalog, error) {
	path := fmt.Sprintf("/businesses/%s/catalogs", businessId)

	catalogs := new(Catalogs)
	err := api.Get(path, nil, catalogs)

	if err != nil {
		return nil, err
	}

	return catalogs.Catalogs, nil

}
