package speedex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shippingAddressService_Query(t *testing.T) {
	addresses, err := client.Services.ShippingAddress.Query(ctx)
	assert.Nil(t, err)
	assert.NotEmpty(t, addresses)
}
