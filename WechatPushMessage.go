package requests

import (
	"fmt"
	"time"
)

// 推送

type PushWechatOption struct {
	Token string                 `json:"token"`
	Body  map[string]interface{} `json:"body"`
}

func (P *PushWechatOption) Push() bool {
	var client ClientOption
	client.Url = "https://api.weixin.qq.com/cgi-bin/message/template/send"
	client.Params = map[string]string{
		"access_token": P.Token,
	}
	client.Body = P.Body
	res := client.ToJson(client.Post())
	fmt.Println(res)
	if res.Get("errcode").Int() != 0 {
		fmt.Println(res.Get("errmsg").String())
		return false
	}
	return true
}

// 模版信息

func (P *PushWechatOption) OutUserTemplate(openid string, Card string, State string, Time string, Remark string) {
	if Card != "888" {
		Remark = ""
	}
	P.Body = map[string]interface{}{
		"touser":      openid,
		"template_id": "1FDV3Bqr_HEApBqsTo128ZF51RfDINI_pr_IuSKtgYw",
		"url":         "https://m.bokao2o.com/store/PDKQG/cardBag",
		"miniprogram": map[string]string{
			"appid":    "wx8f28809a734aad89",
			"pagepath": "pages/account/account?state=query_consumption_records",
		},
		"data": map[string]interface{}{
			"first": map[string]interface{}{
				"value": fmt.Sprintf("您好，您的卡号 %s 账户发生变动:", Card),
				"color": "#FF0000",
			},
			"keyword1": map[string]interface{}{
				"value": fmt.Sprintf("【%s】", State),
				"color": "#003399",
			},
			"keyword2": map[string]interface{}{
				"value": Time,
				"color": "#003399",
			},
			"remark": map[string]interface{}{
				"value": Remark,
				"color": "#FF0900",
			},
		},
	}
}

func (P *PushWechatOption) AdminTemplate(Openid string, Title string, Text string, Time string, Remark string) {
	P.Body = map[string]interface{}{
		"touser":      Openid,
		"template_id": "a8D6JUfvddA-1wdAKsMQrkcBuKtFs8miri_KBQrUYSQ",
		"url":         "https://www.shkqg.com",
		"data": map[string]interface{}{
			"first": map[string]interface{}{
				"value": Title,
				"color": "#FF0000",
			},
			"keyword1": map[string]interface{}{
				"value": Text,
				"color": "#003399",
			},
			"keyword2": map[string]interface{}{
				"value": Time,
				"color": "#003399",
			},
			"remark": map[string]interface{}{
				"value": Remark,
				"color": "#FF0900",
			},
		},
	}
}

func (P *PushWechatOption) UserTemplate(Openid string, Data string, Num string) {
	P.Body = map[string]interface{}{
		"touser":      Openid,
		"template_id": "CqP6DRSZ8VOgB_8BA9Ch3cYdxwpbzgFOvpSvCBSFfxE",
		"url":         fmt.Sprintf("https://www.shkqg.com/message/%s", Data),
		"data": map[string]interface{}{
			"first": map[string]interface{}{
				"value": fmt.Sprintf("工号 %s,  业绩如下：", Num),
				"color": "#FF0000",
			},
			"keyword1": map[string]interface{}{
				"value": time.Now().Format("2006-01") + "01 07:00",
				"color": "#003399",
			},
			"keyword2": map[string]interface{}{
				"value": time.Now().Format("2006-01-02 03:04"),
				"color": "#003399",
			},
			"remark": map[string]interface{}{
				"value": Data,
				"color": "#FF0900",
			},
		},
	}
}

func (P *PushWechatOption) CeShiTemplate(openid string, Card string, State string, Time string, Remark string) {
	if Card != "888" {
		Remark = ""
	}
	P.Body = map[string]interface{}{
		"touser":      openid,
		"template_id": "Jmb15makmeRCSHCIifxWYGApK_7TF0CT2H4lRnjIEko",
		"url":         "https://m.bokao2o.com/store/PDKQG/cardBag",
		"data": map[string]interface{}{
			"first": map[string]interface{}{
				"value": fmt.Sprintf("您好，您的卡号 %s 账户发生变动:", Card),
				"color": "#FF0000",
			},
			"keyword1": map[string]interface{}{
				"value": fmt.Sprintf("【%s】", State),
				"color": "#003399",
			},
			"keyword2": map[string]interface{}{
				"value": Time,
				"color": "#003399",
			},
			"remark": map[string]interface{}{
				"value": Remark,
				"color": "#FF0900",
			},
		},
	}
}
