package topup

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	_ "github.com/stripe/stripe-go/testing"
)

func TestChargeGet(t *testing.T) {
	topup, err := Get("tu_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}
