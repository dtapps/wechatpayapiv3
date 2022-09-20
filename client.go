package wechatpayapiv3

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId            string             // 小程序或者公众号唯一凭证
	AppSecret        string             // 小程序或者公众号唯一凭证密钥
	MchId            string             // 微信支付的商户id
	AesKey           string             // 私钥
	ApiV3            string             // API v3密钥
	MchSslSerialNo   string             // pem 证书号
	MchSslKey        string             // pem key 内容
	ApiGormClientFun golog.ApiClientFun // 日志配置
	Debug            bool               // 日志开关
	ZapLog           *golog.ZapLog      // 日志服务
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	config        struct {
		appId          string // 小程序或者公众号唯一凭证
		appSecret      string // 小程序或者公众号唯一凭证密钥
		mchId          string // 微信支付的商户id
		aesKey         string // 私钥
		apiV3          string // API v3密钥
		mchSslSerialNo string // pem 证书号
		mchSslKey      string // pem key 内容
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.zapLog = config.ZapLog

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.mchId = config.MchId
	c.config.aesKey = config.AesKey
	c.config.apiV3 = config.ApiV3
	c.config.mchSslSerialNo = config.MchSslSerialNo
	c.config.mchSslKey = config.MchSslKey

	c.requestClient = gorequest.NewHttp()

	apiGormClient := config.ApiGormClientFun()
	if apiGormClient != nil {
		c.log.client = apiGormClient
		c.log.status = true
	}

	return c, nil
}
