package requests

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"os"
)

func CreateWechatMenu(token string, menusJson string, defaultJsonFilePath string) string {
	menus := ReadMenuJsonFile(defaultJsonFilePath)

	if menusJson == "moban" {
		data, _ := json.Marshal(menus)
		return string(data)
	}

	Client := ClientOption{
		Url: "https://api.weixin.qq.com/cgi-bin/menu/create",
		Params: map[string]string{
			"access_token": token,
		},
		Headers: nil,
		Body:    menus,
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

func ReadMenuJsonFile(fileName string) (data map[string]interface{}) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	var ok bool
	data, ok = gjson.Parse(string(content)).Value().(map[string]interface{})
	if !ok {
		return
	}
	return
}
