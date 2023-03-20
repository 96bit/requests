package requests

import (
	"fmt"
)

type BoKaLoginConfig struct {
	CustID   string `json:"custId"`
	CompID   string `json:"compid"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Source   string `json:"source"`
}

func (config *BoKaLoginConfig) GetBoKaToken() (token string, shopId string) {
	fmt.Println("开始获取 NEW TOKEN...")
	client := ClientOption{
		Url:    "https://api.bokao2o.com/auth/merchant/v2/user/login",
		Params: nil,
		Headers: map[string]string{
			"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.3",
			"referer":    "https://s3.boka.vc/",
		},
		Body: map[string]interface{}{
			"custId":   config.CustID,
			"compId":   config.CompID,
			"userName": config.UserName,
			"passWord": config.PassWord,
			"source":   config.Source,
		},
	}
	res := client.Post()
	data := client.ToJson(res)
	if data.Get("code").Num == 200 {
		return data.Get("result.token").String(),
			data.Get("result.shopId").String()
	}

	return "", ""
}
