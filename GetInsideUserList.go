package requests

import "errors"

func GetInsideUserList(token string) (users [][]string, err error) {

	client := ClientOption{
		Url:    "https://api.bokao2o.com/s3connect/job/employ/v2/comp/002/getLs?disable=1&sign=UERLUUcjMDAy",
		Params: nil,
		Headers: map[string]string{
			"access_token": token,
			"accesstoken":  token,
			"device_id":    "s3backend",
			"deviceid":     "s3backend",
			"origin":       "https://s3.boka.vc",
			"referer":      "https://s3.boka.vc/home",
		},
		Body: nil,
	}
	res := client.Get()
	data := client.ToJson(res)
	if data.Get("code").Int() != 200 {
		err = errors.New("发生错误")
		return
	}
	for _, value := range data.Get("result").Array() {

		disable := value.Get("disable").Int()
		if disable == 0 {
			number := value.Get("empId").String()
			phone := value.Get("mobile").String()
			users = append(users, []string{number, phone})
		}

	}

	return
}
