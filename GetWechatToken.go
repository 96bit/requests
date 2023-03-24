package requests

import (
	"log"
)

var WECHATOKEN string

type WeChatAccessTokenConfig struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

func (config *WeChatAccessTokenConfig) GetAccessToken() {
	Client := ClientOption{
		Url: "https://api.weixin.qq.com/cgi-bin/stable_token",
		Body: map[string]interface{}{
			"grant_type":    "client_credential",
			"appid":         config.Appid,
			"secret":        config.Secret,
			"force_refresh": false,
		},
	}
	res := Client.ToJson(Client.Post())
	WECHATOKEN = res.Get("access_token").String()

	log.Printf("GetAccessToken: %s", res)
}
