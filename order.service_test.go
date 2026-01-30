package speedex

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func Test_orderService_Create(t *testing.T) {
	addressInfo := gofakeit.Address()
	req := []CreateOrderRequest{
		{
			CustomerOrderNo:          "XY-112-ABC",
			ProductCode:              null.String{},
			ConsigneeName:            gofakeit.Username(),
			ConsigneeCompanyName:     null.String{},
			ConsigneeStateOrProvince: addressInfo.State,
			ConsigneeCity:            addressInfo.City,
			ConsignessArea:           null.String{},
			ConsigneeAddress1:        addressInfo.Address,
			ConsigneeAddress2:        null.String{},
			ConsigneeAddress3:        null.String{},
			ConsigneePostCode:        null.StringFrom(addressInfo.Zip),
			ConsigneePhone:           gofakeit.Phone(),
			ShipperAddressId:         2,
			SignatureService:         "NOSIGNATURE",
			InsuredValue:             null.Float{},
			SizeWeightUnit:           "MET",
			Boxes: []OrderBox{
				{
					No:     1,
					Length: 10,
					Width:  10,
					Height: 10,
					Weight: 1,
					Skus: []OrderBoxSku{
						{
							SKU:         "SKU-1",
							ChineseName: "中文品名-1",
							EnglishName: "English Name-1",
							Quantity:    1,
						},
					},
				},
			},
			AdditionalProperties: nil,
			Notes:                null.String{},
		},
	}
	results, err := client.Services.Order.Create(ctx, req)
	if err == nil {
		assert.Equal(t, "XY-112-ABC", results[0].CustomerNo)
	} else {
		assert.Contains(t, []string{"400:", "500:"}, err.Error()[:4])
	}
}

func Test_orderService_Query(t *testing.T) {
	customerNo := "XY-112-ABC"
	res, err := client.Services.Order.Query(ctx, OrderQueryRequest{
		CustomerNos: customerNo,
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
	if len(res) != 0 {
		assert.Equal(t, customerNo, res[0].CustomerNo)
	}
}

func Test_orderService_Cancel(t *testing.T) {
	results, err := client.Services.Order.Cancel(ctx, CancelOrderRequest{
		OrderNos: []string{"000010202509290000001"},
	})
	assert.Nil(t, err)
	assert.Equal(t, "000010202509290000001", results[0].OrderNo)
}
