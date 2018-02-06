package topup

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /topups APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a topup.
// For more details see https://stripe.com/docs/api#retrieve_topup.
func Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	topup := &stripe.Topup{}
	err := c.B.Call("GET", "/topups/"+id, c.Key, body, commonParams, topup)
	return topup, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
