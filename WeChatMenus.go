package requests

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

func CreateWechatMenu(token string, menusJson string) string {
	menusa := []byte(`{"button":[{"type":"click","name":"预约","key":"yue"},{"type":"view","name":"绑定","url":"https://m.bokao2o.com/bmall/PDKQG/cardBind"},{"name":"会员服务","sub_button":{"list":[{"type":"miniprogram","name":"账户查询","url":"http://mp.weixin.qq.com","appid":"wx8f28809a734aad89","pagepath":"pages/account/account"},{"type":"view","name":"备用查询","url":"https://m.bokao2o.com/bmall/PDKQG/personalCenter"}]}}]}`)

	if menusJson == "moban" {
		return string(menusa)
	}

	menusb := make(map[string]interface{})

	menusaa := []byte(menusJson)
	err := json.Unmarshal(menusaa, &menusb)
	if err != nil {
		fmt.Println(err)
		return "JSON 文件格式错误! "
	}

	fmt.Println(menusb)

	Client := ClientOption{
		Url: "https://api.weixin.qq.com/cgi-bin/menu/create",
		Params: map[string]string{
			"access_token": token,
		},
		Headers: nil,
		Body:    menusb,
	}

	res := Client.ToJson(Client.Post())
	return res.String()
}

func QueryWechatMenu(token string) gjson.Result {
	Client := ClientOption{
		Url: "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info",
		Params: map[string]string{
			"access_token": token,
		},
		Headers: nil,
		Body:    nil,
	}

	res := Client.ToJson(Client.Get())
	return res
}
