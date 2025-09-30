package speedex

import (
	"context"

	"github.com/hiscaler/speedex-go/entity"
)

// productService 产品服务
type productService service

func (s productService) Query(ctx context.Context) ([]entity.Product, error) {
	var res []entity.Product
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetResult(&res).
		Get("/external/products")
	if err != nil {
		return nil, err
	}

	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}
