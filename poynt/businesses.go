package poynt

import (
	"fmt"
)

type Business interface{}

// GetBusiness returns a single catalog
func (api *PoyntApi) GetBusiness(businessId string) (*Business, error) {
	path := fmt.Sprintf("/businesses/%s", businessId)

	business := new(Business)
	err := api.Get(path, nil, business)

	if err != nil {
		return nil, err
	}

	fmt.Println("wtf", business)

	return business, nil
}
