package entity

import "gopkg.in/guregu/null.v4"

// ShippingAddress 发货地址
type ShippingAddress struct {
	Id                     int         `json:"id"`                     // 地址 ID
	ShipperName            string      `json:"shipperName"`            // 发货人姓名
	ShipperCompanyName     null.String `json:"shipperCompanyName"`     // 发货公司名称
	ShipperCountryCode     string      `json:"shipperCountryCode"`     // 发货国家代码
	ShipperStateOrProvince string      `json:"shipperStateOrProvince"` // 州/省
	ShipperCity            string      `json:"shipperCity"`            // 城市
	ShipperArea            null.String `json:"shipperArea"`            // 地区
	ShipperAddress1        string      `json:"shipperAddress1"`        // 地址1
	ShipperAddress2        null.String `json:"shipperAddress2"`        // 地址2
	ShipperAddress3        null.String `json:"shipperAddress3"`        // 地址3
	ShipperPostCode        string      `json:"shipperPostCode"`        // 邮政编码
	ShipperPhone           string      `json:"shipperPhone"`           // 联系电话
	CreatedAt              null.Time   `json:"createdAt"`              // 创建时间
	UpdatedAt              null.Time   `json:"updatedAt"`              // 更新时间
	DeletedAt              null.Time   `json:"deletedAt"`              // 删除时间
	Active                 bool        `json:"active"`                 // 是否启用
	UserId                 int         `json:"userId"`                 // 用户 ID

}
