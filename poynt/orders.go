package poynt

import (
	"fmt"
)

type OrderContext struct {
	BusinessId     string `json:"businessId"`
	StoreId        string `json:"storeId"`
	DeviceId       string `json:"storeDeviceId"`
	EmployeeUserId int64  `json:"employeeUserId"`
	Source         string `json:"source"`
	SourceApp      string `json:"sourceApp"`
}

type OrderAmounts struct {
	Currency      string `json:"currency"`
	DiscountTotal int64  `json:"discountTotal"`
	FeeTotal      int64  `json:"feeTotal"`
	NetTotal      int64  `json:"netTotal"`
	SubTotal      int64  `json:"subTotal"`
	TaxTotal      int64  `json:"taxTotal"`
}

type OrderDiscount struct {
	Amount     int64   `json:"amount"`
	CustomName string  `json:"customName"`
	Fixed      int64   `json:"fixed"`
	Id         string  `json:"id"`
	Percentage float64 `json:"percentage"`
	Processor  string  `json:"processor"`
	Provider   int64   `json:"provider"`
	Type       string  `json:"type"`
}

type OrderFee struct {
	Amount     int64   `json:"amount"`
	CustomName string  `json:"customName"`
	Fixed      int64   `json:"fixed"`
	Id         string  `json:"id"`
	Percentage float64 `json:"percentage"`
	Processor  string  `json:"processor"`
	Provider   int64   `json:"provider"`
	Type       string  `json:"type"`
}

type OrderTax struct {
	Amount       int64  `json:"amount"`
	CatalogLevel bool   `json:"catalogLevel"`
	Id           string `json:"id"`
	TaxExempted  bool   `json:"taxExempted"`
	Type         string `json:"type"`
}

type OrderSelectedVariant interface{}

type OrderItem struct {
	Name             string                  `json:"name"`
	UnitPrice        int64                   `json:"unitPrice"`
	ProductId        string                  `json:"productId"`
	Sku              string                  `json:"sku"`
	Discount         int64                   `json:"discount"`
	Fee              int64                   `json:"fee"`
	Tax              int64                   `json:"tax"`
	Status           string                  `json:"status"`
	UnitOfMeasure    string                  `json:"unitOfMeasure"`
	TaxExempted      bool                    `json:"taxExempted"`
	Quantity         int64                   `json:"quantity"`
	Notes            string                  `json:"notes"`
	Discounts        []*OrderDiscount        `json:"discounts,omitempty"`
	Fees             []*OrderFee             `json:"fees,omitempty"`
	Taxes            []*OrderTax             `json:"taxes,omitempty"`
	SelectedVariants []*OrderSelectedVariant `json:"selectedVariants,omitempty"`
}

type OrderStatuses struct {
	FulfillmentStatus        string `json:"fulfillmentStatus"`
	Status                   string `json:"status"`
	TransactionStatusSummary string `json:"transactionStatusSummary"`
}

type Order struct {
	Context     *OrderContext    `json:"context"`
	Statuses    *OrderStatuses   `json:"statuses"`
	Amounts     *OrderAmounts    `json:"amounts"`
	OrderNumber int64            `json:"orderNumber"`
	Id          string           `json:"id"`
	TaxExempted bool             `json:"taxExempted"`
	Notes       string           `json:"notes"`
	Discounts   []*OrderDiscount `json:"discounts"`
	Fees        []*OrderFee      `json:"fees"`
	Items       []*OrderItem     `json:"items"`
}

type ResponseOrder struct{}

// Creates single order
func (self *PoyntApi) CreateOrder(businessId string, order *Order) error {
	path := fmt.Sprintf("/businesses/%s/orders", businessId)

	newOrder := new(ResponseOrder)
	err := self.Post(path, order, newOrder)

	if err != nil {
		return err
	}

	return nil
}
