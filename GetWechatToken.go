package requests

import "fmt"

type WeChatAccessTokenConfig struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

func (config *WeChatAccessTokenConfig) GetAccessToken() string {
	Client := ClientOption{
		Url: "https://api.weixin.qq.com/cgi-bin/token",
		Params: map[string]string{
			"grant_type": "client_credential",
			"appid":      config.Appid,
			"secret":     config.Secret,
		},
	}
	res := Client.ToJson(Client.Get())
	Token := res.Get("access_token").String()
	if Token == "" {
		fmt.Println(fmt.Sprintf("GetAccessToken: %s", Token))
	}
	fmt.Println(fmt.Sprintf("GetAccessToken: %s", res))
	return Token
}
