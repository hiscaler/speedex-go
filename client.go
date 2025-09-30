package speedex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/speedex-go/config"
	"github.com/hiscaler/speedex-go/entity"
)

const (
	OK              = 200 // 无错误
	BadRequestError = 400 // 请求错误
	InvalidToken    = 401 // 无效的 Token
	InternalError   = 500 // 内部服务器错误
)

const (
	Version   = "0.0.1"
	userAgent = "SpeedEx API Client-Golang/" + Version + " (https://github.com/hiscaler/speedex-go)"
)

const (
	ProdBaseUrl = "https://speedex.net.cn/apiv1"
	TestBaseUrl = "https://test.speedex.net.cn/apiv1"
)

type Client struct {
	config      *config.Config // 配置
	httpClient  *resty.Client  // Resty Client
	accessToken string         // AccessToken
	Services    services       // API Services
}

func NewClient(ctx context.Context, cfg config.Config) *Client {
	logger := log.New(os.Stdout, "[ Client ] ", log.LstdFlags|log.Llongfile)
	speedExClient := &Client{
		config: &cfg,
	}
	baseUrl := ProdBaseUrl
	if cfg.Env != entity.Prod {
		baseUrl = TestBaseUrl
	}
	httpClient := resty.New().
		SetDebug(cfg.Debug).
		SetBaseURL(baseUrl).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	httpClient.SetTimeout(time.Duration(cfg.Timeout) * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			if speedExClient.accessToken == "" {
				err := speedExClient.getAccessToken(ctx)
				if err != nil {
					return err
				}
			}
			client.SetAuthToken(speedExClient.accessToken)
			return nil
		}).
		SetRetryCount(2).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)

	speedExClient.httpClient = httpClient

	xService := service{
		config:     &cfg,
		logger:     logger,
		httpClient: speedExClient.httpClient,
	}
	speedExClient.Services = services{
		Order:           (orderService)(xService),
		Product:         (productService)(xService),
		ScanForm:        (scanFormService)(xService),
		ShippingAddress: (shippingAddressService)(xService),
	}
	return speedExClient
}

// accessToken 获取 Access Token 值
// https://docs.speedex.net.cn/315079583e0
func (c *Client) getAccessToken(ctx context.Context) error {
	if c.accessToken != "" {
		return nil
	}

	var result entity.UserInformation
	baseUrl := ProdBaseUrl
	if c.config.Env != entity.Prod {
		baseUrl = TestBaseUrl
	}
	httpClient := resty.New().
		SetDebug(c.config.Debug).
		SetBaseURL(baseUrl).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	resp, err := httpClient.R().
		SetContext(ctx).
		SetBody(map[string]string{
			"account":  c.config.Account,
			"password": c.config.Password,
		}).
		SetResult(&result).
		Post("/external/login")
	if err = recheckError(resp, err); err != nil {
		return err
	}
	c.accessToken = result.Token
	return nil
}

type NormalResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

// errorWrap 错误包装
func errorWrap(code int, message string) error {
	if code == OK || code == 0 {
		return nil
	}

	switch code {
	case InvalidToken:
		message = "无效的 Token"
	default:
		if code == InternalError {
			if message == "" {
				message = "内部服务器错误，请联系【闪派国际】客服人员"
			}
		} else {
			message = strings.TrimSpace(message)
			if message == "" {
				message = "Unknown error"
			}
		}
	}
	return fmt.Errorf("%d: %s", code, message)
}

func invalidInput(e error) error {
	var errs validation.Errors
	if !errors.As(e, &errs) {
		return e
	}

	if len(errs) == 0 {
		return nil
	}

	fields := make([]string, 0)
	messages := make([]string, 0)
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)

	for _, field := range fields {
		e1 := errs[field]
		if e1 == nil {
			continue
		}

		var errObj validation.ErrorObject
		if errors.As(e1, &errObj) {
			e1 = errObj
		} else {
			var errs1 validation.Errors
			if errors.As(e1, &errs1) {
				e1 = invalidInput(errs1)
				if e1 == nil {
					continue
				}
			}
		}

		messages = append(messages, e1.Error())
	}
	return errors.New(strings.Join(messages, "; "))
}

func recheckError(resp *resty.Response, e error) error {
	if e != nil {
		if errors.Is(e, http.ErrHandlerTimeout) {
			return errorWrap(http.StatusRequestTimeout, e.Error())
		}
		return e
	}

	if resp.IsError() {
		var normalResponse NormalResponse
		err := json.Unmarshal(resp.Body(), &normalResponse)
		if err != nil {
			return err
		}
		return errorWrap(resp.StatusCode(), normalResponse.Message)
	}

	return nil
}
