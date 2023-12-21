package wechatpayapiv3

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type PayTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

type PayTransactionsJsapiResult struct {
	Result PayTransactionsJsapiResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newPayTransactionsJsapiResult(result PayTransactionsJsapiResponse, body []byte, http gorequest.Response) *PayTransactionsJsapiResult {
	return &PayTransactionsJsapiResult{Result: result, Body: body, Http: http}
}

// PayTransactionsJsapi JSAPI下单
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_1.shtml
func (c *Client) PayTransactionsJsapi(ctx context.Context, notMustParams ...gorequest.Params) (*PayTransactionsJsapiResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/transactions/jsapi", params, http.MethodPost, true)
	if err != nil {
		return newPayTransactionsJsapiResult(PayTransactionsJsapiResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PayTransactionsJsapiResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPayTransactionsJsapiResult(response, request.ResponseBody, request), err
}
