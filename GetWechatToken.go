package requests

import (
	"errors"
	"log"
)

type WeChatAccessTokenConfig struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

func (config *WeChatAccessTokenConfig) GetAccessToken() (token string, err error) {
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
	Token := res.Get("access_token").String()
	if Token == "" {
		return "", errors.New(res.String())
	}
	log.Printf("GetAccessToken: %s", res)

	return Token, nil
}
