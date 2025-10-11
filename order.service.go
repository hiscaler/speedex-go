package speedex

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/speedex-go/entity"
	"gopkg.in/guregu/null.v4"
)

// 订单服务
type orderService service

// OrderBox 订单箱子
type OrderBox struct {
	No     int           `json:"no"`     // 箱子编号
	Length float64       `json:"length"` // 箱子长度
	Width  float64       `json:"width"`  // 箱子宽度
	Height float64       `json:"height"` // 箱子高度
	Weight float64       `json:"weight"` // 箱子重量
	Skus   []OrderBoxSku `json:"skus"`   // SKU 列表
}

type OrderBoxSku struct {
	SKU         string `json:"sku"`         // SKU 编码
	ChineseName string `json:"chineseName"` // 中文品名
	EnglishName string `json:"englishName"` // 英文品名
	Quantity    int    `json:"quantity"`    // 数量
}

func (m OrderBox) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Length,
			validation.Required.Error("长不能为空"),
			validation.Min(0.01).Error("长不能小于 {{.min}}"),
			validation.Max(999999.99).Error("长不能大于 {{.max}}"),
		),
		validation.Field(&m.Width,
			validation.Required.Error("宽不能为空"),
			validation.Min(0.01).Error("宽不能小于 {{.min}}"),
			validation.Max(999999.99).Error("宽不能大于 {{.max}}"),
		),
		validation.Field(&m.Height,
			validation.Required.Error("高不能为空"),
			validation.Min(0.01).Error("高不能小于 {{.min}}"),
			validation.Max(999999.99).Error("高不能大于 {{.max}}"),
		),
		validation.Field(&m.Weight,
			validation.Required.Error("重量不能为空"),
			validation.Min(0.01).Error("重量不能小于 {{.min}}"),
			validation.Max(999999.99).Error("重量不能大于 {{.max}}"),
		),
	)
}

type CreateOrderRequest struct {
	CustomerOrderNo          string      `json:"customerOrderNo"`                // 客户单号
	ProductCode              null.String `json:"productCode,omitempty"`          // 产品代码
	ConsigneeName            string      `json:"consigneeName"`                  // 收件人姓名
	ConsigneeCompanyName     null.String `json:"consigneeCompanyName,omitempty"` // 收件人公司名称
	ConsigneeStateOrProvince string      `json:"consigneeStateOrProvince"`       // 收件人省/州
	ConsigneeCity            string      `json:"consigneeCity"`                  // 收件人城市
	ConsignessArea           null.String `json:"consignessArea,omitempty"`       // 收件人地区
	ConsigneeAddress1        string      `json:"consigneeAddress1"`              // 收件人地址1
	ConsigneeAddress2        null.String `json:"consigneeAddress2,omitempty"`    // 收件人地址2
	ConsigneeAddress3        null.String `json:"consigneeAddress3,omitempty"`    // 收件人地址3
	ConsigneePostCode        null.String `json:"consigneePostCode,omitempty"`    // 收件人邮编
	ConsigneePhone           string      `json:"consigneePhone"`                 // 收件人电话
	ShipperAddressId         int         `json:"shipperAddressId"`               // 发件人地址 ID
	SignatureService         string      `json:"signatureService"`               // 签名服务（签名服务类型：ASS（成人签名）、SSF（普通签名）、NOSIGNATURE（无签名服务））
	InsuredValue             null.Float  `json:"insuredValue"`                   // 保价价值，最多 10 位数字，保留 2 位小数
	SizeWeightUnit           string      `json:"sizeWeightUnit"`                 // 重量尺寸单位：公制(MET)或英制(IMP)
	Boxes                    []OrderBox  `json:"boxes"`                          // 订单箱子列表
	AdditionalProperties     any         `json:"additionalProperties,omitempty"` // 附加属性
	Notes                    null.String `json:"notes,omitempty"`                // 备注
}

func (m CreateOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.CustomerOrderNo,
			validation.Required.Error("客户单号不能为空"),
			validation.Length(1, 255).Error("客户单号不能超过 {{.max}} 个字符"),
		),
		validation.Field(&m.ConsigneeName,
			validation.Required.Error("收件人不能为空"),
			validation.Length(3, 255).Error("收件人长度必须在 {{.min}} ~ {{.max}} 个字符"),
		),
		validation.Field(&m.ConsigneeCompanyName, validation.When(m.ConsigneeCompanyName.Valid, validation.Length(1, 255).Error("收件人公司不能超过 {{.max}} 个字符"))),
		validation.Field(&m.ConsigneeStateOrProvince, validation.Required.Error("收件人州不能为空")),
		validation.Field(&m.ConsigneeCity, validation.Required.Error("收件人城市不能为空")),
		//validation.Field(&m.ConsignessArea, validation.Required.Error("收件人地区不能为空")),
		validation.Field(&m.ConsigneeAddress1,
			validation.Required.Error("收件人地址1不能为空"),
			validation.Length(1, 255).Error("收件人地址1长度不能超过 {{.max}} 个字符"),
		),
		validation.Field(&m.ConsigneeAddress2,
			validation.When(m.ConsigneeAddress2.Valid, validation.Length(1, 255).Error("收件人地址2长度不能超过 {{.max}} 个字符")),
		),
		validation.Field(&m.ConsigneeAddress3,
			validation.When(m.ConsigneeAddress3.Valid, validation.Length(1, 255).Error("收件人地址3长度不能超过 {{.max}} 个字符")),
		),
		validation.Field(&m.ConsigneePostCode, validation.Required.Error("收件人邮编不能为空")),
		validation.Field(&m.ConsigneePhone,
			validation.Required.Error("收件人电话不能为空"),
			validation.Length(10, 15).Error("收件人电话长度必须在 {{.min}} ~ {{.max}} 个字符"),
		),
		validation.Field(&m.ShipperAddressId, validation.Required.Error("发件人地址不能为空")),
		validation.Field(&m.SignatureService,
			validation.When(m.SignatureService != "", validation.In("ASS", "SSF", "NOSIGNATURE").Error("签名服务参数错误")),
		),
		validation.Field(&m.SizeWeightUnit, validation.In("MET", "IMP").Error("重量尺寸单位类型参数错误")),
		validation.Field(&m.Boxes, validation.Required.Error("包裹信息不能为空")),
		validation.Field(&m.Notes, validation.When(m.Notes.Valid, validation.Length(1, 255).Error("备注不能超过 {{.max}} 个字符"))),
	)
}

type CreateOrderResult struct {
	CustomerNo string         `json:"customerNo"` // 客户订单号
	Orders     []entity.Order `json:"orders"`     // 订单列表
}

// Create 异步创建订单
// https://docs.speedex.net.cn/315079587e0
func (s orderService) Create(ctx context.Context, requests []CreateOrderRequest) ([]CreateOrderResult, error) {
	for _, req := range requests {
		if err := req.Validate(); err != nil {
			return nil, invalidInput(err)
		}
	}

	var res []CreateOrderResult
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(requests).
		SetResult(&res).
		Post("/external/orders/async")
	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}

// RetryCreate 重新下单
func (s orderService) RetryCreate(ctx context.Context, orderNumbers ...string) error {
	if len(orderNumbers) == 0 {
		return errors.New("订单号不能为空")
	}

	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(map[string][]string{"orderNos": orderNumbers}).
		Post("/external/orders/reload")
	return recheckError(resp, err)
}

type OrderQueryRequest struct {
	CustomerNos string `json:"customerNos"` // 订单号
}

func (m OrderQueryRequest) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.CustomerNos, validation.Required.Error("订单号不能为空")),
	)
}

// Query 根据查询条件筛选符合条件的订单列表数据
// https://docs.speedex.net.cn/315079588e0
func (s orderService) Query(ctx context.Context, req OrderQueryRequest) ([]entity.Order, error) {
	if err := req.validate(); err != nil {
		return nil, invalidInput(err)
	}

	var res []entity.Order
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"customerNos": req.CustomerNos,
		}).
		SetResult(&res).
		Get("/external/orders")
	if err != nil {
		return nil, err
	}

	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}

type CancelOrderRequest struct {
	OrderNos string `json:"orderNos"` // 订单号
}

func (m CancelOrderRequest) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderNos, validation.Required.Error("订单号不能为空")),
	)
}

type OrderCancelResult struct {
	OrderNo    string      `json:"orderNo"`    // 订单号
	FailReason null.String `json:"failReason"` // 失败原因
}

// Cancel 取消订单
// https://docs.speedex.net.cn/315514901e0
func (s orderService) Cancel(ctx context.Context, req CancelOrderRequest) ([]OrderCancelResult, error) {
	if err := req.validate(); err != nil {
		return nil, invalidInput(err)
	}

	var res []OrderCancelResult
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		Delete("/external/orders")
	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}

type OrderEstimateRequest = CreateOrderRequest

// Estimate 订单费用试算
func (s orderService) Estimate(ctx context.Context, req OrderEstimateRequest) (res entity.OrderEstimateResult, err error) {
	if err = req.Validate(); err != nil {
		return res, invalidInput(err)
	}

	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		Post("/external/orders/estimate")
	if err = recheckError(resp, err); err != nil {
		return res, err
	}
	return res, nil
}
