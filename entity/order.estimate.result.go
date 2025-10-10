package entity

import "gopkg.in/guregu/null.v4"

// OrderEstimateResult 订单费用试算
type OrderEstimateResult struct {
	CustomerOrderNo string                      `json:"customerOrderNo"` // 客户单号
	ProductPrices   []OrderEstimateProductPrice `json:"productPrices"`   // 产品报价列表
}

// OrderEstimateProductPrice 订单费用试算-商品费用
type OrderEstimateProductPrice struct {
	ProductCode         string      `json:"productCode"`
	ProductId           int         `json:"productId"`
	ProductName         string      `json:"productName"`
	Currency            string      `json:"currency"`
	Price               string      `json:"price"`
	SellingFreightPrice string      `json:"sellingFreightPrice"`
	SellingMiscFeePrice string      `json:"sellingMiscFeePrice"`
	FailReason          null.String `json:"failReason"`
}
