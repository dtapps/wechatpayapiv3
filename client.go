package wechatpayapiv3

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

type ConfigClient struct {
	AppId          string           // 小程序或者公众号唯一凭证
	AppSecret      string           // 小程序或者公众号唯一凭证密钥
	MchId          string           // 微信支付的商户id
	AesKey         string           // 私钥
	ApiV3          string           // API v3密钥
	MchSslSerialNo string           // pem 证书号
	MchSslKey      string           // pem key 内容
	GormClient     *dorm.GormClient // 日志数据库
	LogClient      *golog.ZapLog    // 日志驱动
	LogDebug       bool             // 日志开关
}

// Client 微信支付直连商户
type Client struct {
	requestClient *gorequest.App   // 请求服务
	logClient     *golog.ApiClient // 日志服务
	config        *ConfigClient    // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.requestClient = gorequest.NewHttp()

	if c.config.GormClient.Db != nil {
		c.logClient, err = golog.NewApiClient(&golog.ApiClientConfig{
			GormClient: c.config.GormClient,
			TableName:  logTable,
			LogClient:  c.config.LogClient,
			LogDebug:   c.config.LogDebug,
		})
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
