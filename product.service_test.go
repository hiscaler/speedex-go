package speedex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productService_Query(t *testing.T) {
	products, err := client.Services.Product.Query(ctx)
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
}
