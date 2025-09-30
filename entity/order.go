package entity

import "gopkg.in/guregu/null.v4"

// Order 订单
type Order struct {
	OrderNo      string      `json:"orderNo"`      // 订单号
	CustomerNo   string      `json:"customerNo"`   // 客户编号
	TrackingNo   null.String `json:"trackingNo"`   // 物流单号
	SellingPrice string      `json:"sellingPrice"` // 销售价格
	Status       string      `json:"status"`       // 订单状态
	FailReason   null.String `json:"failReason"`   // 失败原因
	LabelUrl     null.String `json:"labelUrl"`     // 面单 URL
	UpdatedAt    null.Time   `json:"updatedAt"`    // 更新时间
}
