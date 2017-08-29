package poynt

type TransactionAmounts struct {
	CashbackAmount    int    `json:"cashbackAmount"`
	Currency          string `json:"currency"`
	OrderAmount       int    `json:"orderAmount"`
	TipAmount         int    `json:"tipAmount"`
	TransactionAmount int    `json:"transactionAmount"`
}

type TransactionContext struct {
	BusinessId            string `json:"businessId"`
	BusinessType          string `json:"businessType"`
	EmployeeUserId        string `json:"employeeUserId"`
	Mcc                   string `json:"mcc"`
	Source                string `json:"source"`
	SourceApp             string `json:"sourceApp"`
	StoreAddressCity      string `json:"storeAddressCity"`
	StoreAddressTerritory string `json:"storeAddressTerritory"`
	StoreDeviceId         string `json:"storeDeviceId"`
	StoreId               string `json:"storeId"`
	StoreTimezone         string `json:"storeTimezone"`
	TransmissionAtLocal   string `json:"transmissionAtLocal"`
}

type TransactionFundingSourceCard struct {
	CardHolderFirstName string `json:"cardHolderFirstName"`
	CardHolderFullName  string `json:"cardHolderFullName"`
	CardHolderLastName  string `json:"cardHolderLastName"`
	Encrypted           bool   `json:"encrypted"`
	ExpirationDate      int    `json:"expirationDate"`
	ExpirationMonth     int    `json:"expirationMonth"`
	ExpirationYear      int    `json:"expirationYear"`
	Id                  int    `json:"id"`
	NumberFirst6        string `json:"numberFirst6"`
	NumberLast4         string `json:"numberLast4"`
	NumberMasked        string `json:"numberMasked"`
	ServiceCode         string `json:"serviceCode"`
	Type                string `json:"type"`
}

type TransactionFundingSourceEmvData interface{}

type TransactionFundingSourceEntryDetails interface{}

type TransactionFundingSource struct {
	Card         *TransactionFundingSourceCard         `json:"card"`
	Debit        bool                                  `json:"debit"`
	EmvData      *TransactionFundingSourceEmvData      `json:"emvData"`
	EntryDetails *TransactionFundingSourceEntryDetails `json:"entryDetails"`
	Type         string                                `json:"type"`
}

type TransactionPoyntLoyalty interface{}
type TransactionProcessorResponse interface{}
type TransactionReference struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Transaction struct {
	Action            string                        `json:"action"`
	Amounts           *TransactionAmounts           `json:"amounts"`
	AuthOnly          bool                          `json:"authAonly"`
	Context           *TransactionContext           `json:"transactionContext"`
	CreatedAt         string                        `json:"createdAt"`
	CustomerUserId    string                        `json:"customerUserId"`
	FundingSource     *TransactionFundingSource     `json:"fundingSource"`
	Id                string                        `json:"id"`
	PoyntLoyalty      *TransactionPoyntLoyalty      `json:"poyntLoyalty"`
	ProcessorResponse *TransactionProcessorResponse `json:"transactionProcessorResponse"`
	References        []*TransactionReference       `json:"references"`
	SignatureCaptured string                        `json:"signatureCaptured"`
	Status            string                        `json:"status"`
	UpdatedAt         string                        `json:"updatedAt"`
}
