package wechatpayapiv3

func (c *Client) ConfigApp(appId, appSecret string) *Client {
	c.config.appId = appId
	c.config.appSecret = appSecret
	return c
}
