package speedex

import (
	"context"

	"github.com/hiscaler/speedex-go/entity"
)

// shippingAddressService 发货地址服务
type shippingAddressService service

// Query 发货地址查询
func (s shippingAddressService) Query(ctx context.Context) ([]entity.ShippingAddress, error) {
	var res []entity.ShippingAddress
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetResult(&res).
		Get("/external/shipper-addresses")
	if err != nil {
		return nil, err
	}

	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}
