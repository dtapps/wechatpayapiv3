package wechatpayapiv3

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, method string, commonParams bool) (gorequest.Response, error) {

	// 公共参数
	if method == http.MethodPost {
		if commonParams == true {
			params["appid"] = c.GetAppId()
			params["mchid"] = c.GetMchId()
		}
	}

	// 认证
	authorization, err := c.authorization(method, params, url)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置JSON格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 设置头部
	client.SetHeader("Authorization", "WECHATPAY2-SHA256-RSA2048 "+authorization)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Accept-Language", "zh-CN")
	if url == "https://api.mch.weixin.qq.com/v3/merchant-service/complaints-v2" {
		client.SetHeader("Wechatpay-Serial", c.GetMchSslSerialNo())
	}

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.GormClient.Db != nil {
		go c.logClient.GormMiddleware(ctx, request, Version)
	}

	return request, err
}
