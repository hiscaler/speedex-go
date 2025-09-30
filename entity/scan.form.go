package entity

// ScanForm 交运单
type ScanForm struct {
	TrackingNos    []string `json:"trackingNos"`    // 跟踪号列表
	ProductName    string   `json:"productName"`    // 产品名称
	ProductCode    string   `json:"productcode"`    // 产品代码
	ScanFormUrl    string   `json:"scanFormUrl"`    // ScanForm 单的下载链接
	ScanFormNo     string   `json:"scanFormNo"`     // ScanForm 编号
	TotalTickets   int      `json:"totalTickets"`   // 总票数
	TotalWeight    int      `json:"totalWeight"`    // 总重量
	SizeWeightUnit string   `json:"sizeWeightUnit"` // 重量单位
	CreatedAt      string   `json:"createdAt"`      // 创建时间
}
