package poynt

import "fmt"

type PaymentResponse interface{}

type PaymentData struct {
	Action         string `json:"action"`
	PurchaseAmount int64  `json:"purchaseAmount"`
	TipAmount      int64  `json:"tipAmount"`
	Currency       string `json:"currency"`
	ReferenceId    string `json:"referenceId"`
	OrderId        string `json:"orderId"`
	CallbackUrl    string `json:"callbackUrl"`
	MultiTender    bool   `json:"multiTender"`
}

type PaymentBody struct {
	TTL        int64  `json:"ttl"`
	BusinessId string `json:"businessId"`
	StoreId    string `json:"storeId"`
	DeviceId   string `json:"storeDeviceId"`
	Data       string `json:"data"`
}

// Makes a poynt cloud request to trigger payment on a device
// Creates single order
func (self *PoyntApi) TriggerPayment(body *PaymentBody, data *PaymentData) error {
	data.MultiTender = true

	path := fmt.Sprintf("/cloudMessages")

	resp := new(PaymentResponse)

	dataString, errStringify := Stringify(data)

	if errStringify != nil {
		return errStringify
	}

	body.Data = dataString
	err := self.Post(path, body, resp)

	if err != nil {
		return err
	}

	return nil
}
