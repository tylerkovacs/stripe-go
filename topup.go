package stripe

type TopupParams struct {
	Params `form:"*"`
	Source *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
}

// SetSource adds valid sources to a TopupParams object,
// returning an error for unsupported sources.
func (p *TopupParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
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
