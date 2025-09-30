package speedex

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/speedex-go/entity"
)

// scanFormService 交运单服务
type scanFormService service

type ScanFormCreateRequest struct {
	TrackingNos []string `json:"trackingNos"` // 追踪编号列表
}

func (m ScanFormCreateRequest) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TrackingNos, validation.Required.Error("追踪编号不能为空")),
	)
}

// Create 创建交运清单
// https://docs.speedex.net.cn/315079589e0
func (s scanFormService) Create(ctx context.Context, req ScanFormCreateRequest) (scanForm entity.ScanForm, err error) {
	if err = req.validate(); err != nil {
		return scanForm, invalidInput(err)
	}

	var res entity.ScanForm
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		Post("/external/scanform")
	if err != nil {
		return scanForm, err
	}

	if err = recheckError(resp, err); err != nil {
		return scanForm, err
	}
	return res, nil

}

type ScanFormQueryRequest struct {
	Page        int    `json:"page"`        // 页码
	PageSize    int    `json:"pageSize"`    // 每页数量
	TrackingNos string `json:"trackingNos"` // 用逗号分隔的追踪编号字符串，例如：123,456
}

func (m ScanFormQueryRequest) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TrackingNos, validation.Required.Error("追踪编号不能为空")),
	)
}

// Query 获取历史 scanform 列表
// https://docs.speedex.net.cn/315079590e0
func (s scanFormService) Query(ctx context.Context, req ScanFormQueryRequest) ([]entity.ScanForm, error) {
	if err := req.validate(); err != nil {
		return nil, invalidInput(err)
	}

	var res []entity.ScanForm
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetQueryString(fmt.Sprintf("page=%d&pageSize=%d&trackingNos=%s", req.Page, req.PageSize, req.TrackingNos)).
		SetResult(&res).
		Get("/external/scanforms")
	if err != nil {
		return nil, err
	}

	if err = recheckError(resp, err); err != nil {
		return nil, err
	}
	return res, nil
}
