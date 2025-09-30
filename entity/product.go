package entity

// Product 产品信息
type Product struct {
	Id      int    `json:"id"`      // 产品ID
	Name    string `json:"name"`    // 产品类型（如 Ground）
	Type    string `json:"type"`    // 产品类型
	Code    string `json:"code"`    // 产品编码
	Carrier string `json:"carrier"` // 承运商名称（如 USPS）
}
