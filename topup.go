package stripe

type TopupParams struct {
	Params    `form:"*"`
	Amount    uint64        `form:"amount"`
	Currency  Currency      `form:"currency"`
	Desc      string        `json:"description"`
	Source    *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	Statement string        `json:"statement_descriptor"`
}

// SetSource adds valid sources to a TopupParams object,
// returning an error for unsupported sources.
func (p *TopupParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// TopupListParams is the set of parameters that can be used when listing topups.
// For more details see https://stripe.com/docs/api#list_topups.
type TopupListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Topup is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#topups.
type Topup struct {
	Amount                   uint64         `json:"amount"`
	ArrivalDate              int64          `json:"arrival_date"`
	Created                  int64          `json:"created"`
	Currency                 Currency       `json:"currency"`
	Desc                     string         `json:"description"`
	ExpectedAvailabilityDate int64          `json:"expected_availability_date"`
	FailCode                 string         `json:"failure_code"`
	FailMsg                  string         `json:"failure_message"`
	ID                       string         `json:"id"`
	Live                     bool           `json:"livemode"`
	Source                   *PaymentSource `json:"source"`
	Statement                string         `json:"statement_descriptor"`
	Status                   string         `json:"status"`
	Tx                       *Transaction   `json:"balance_transaction"`
}
