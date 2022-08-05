package wechatpayapiv3

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	AppId          string            // 小程序或者公众号唯一凭证
	AppSecret      string            // 小程序或者公众号唯一凭证密钥
	MchId          string            // 微信支付的商户id
	AesKey         string            // 私钥
	ApiV3          string            // API v3密钥
	MchSslSerialNo string            // pem 证书号
	MchSslKey      string            // pem key 内容
	MongoDb        *dorm.MongoClient // 日志数据库
	PgsqlDb        *gorm.DB          // 日志数据库
	DatabaseName   string            // 库名
}

// Client 微信支付直连商户
type Client struct {
	client *gorequest.App   // 请求客户端
	log    *golog.ApiClient // 日志服务
	config *ConfigClient    // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.client = gorequest.NewHttp()

	if c.config.PgsqlDb != nil {
		c.log, err = golog.NewApiClient(
			golog.WithGormClient(c.config.PgsqlDb),
			golog.WithTableName(logTable),
		)
		if err != nil {
			return nil, err
		}
	}
	if c.config.MongoDb != nil {
		c.log, err = golog.NewApiClient(
			golog.WithMongoClient(c.config.MongoDb),
			golog.WithDatabaseName(c.config.DatabaseName),
			golog.WithCollectionName(logTable),
		)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
