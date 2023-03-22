package requests

import (
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)

type UsersCourse struct {
	Course map[string]UserCourse
}

type UserCourse struct {
	User    string
	Total   int64
	Tickets []UserDetail
}
type UserDetail struct {
	Date            string
	Name            string
	Amount          int64
	ActualAmount    int64
	MethodPayment   string
	ShareProportion float64
	Detail          string
}

func (user *UsersCourse) GetCourses(data gjson.Result) {
	result := data.Get("result").Array()
	user.Course = make(map[string]UserCourse)

	for _, k := range result {
		userId, course := user.GetCourse(k)
		user.Course[userId] = course
	}
	return
}

func (user *UsersCourse) GetCourse(data gjson.Result) (personId string, userCourse UserCourse) {
	if data.Get("code").Int() != 200 {
		return
	}
	result := data.Get("result.0").Array()
	// 获取工号

	userCourse.User = personId
	for k := range result {
		if result[k].Get("action_id").Int() == 0 && result[k].Get("person_id").Int() != 0 {
			if personId == "" {
				personId = result[k].Get("person_id").String()
			}
			//a := result[k].Get("amt").Num
			var userDetail UserDetail

			userDetail.Amount = result[k].Get("amt").Int()
			userDetail.ActualAmount = result[k].Get("amt3").Int()
			userDetail.Date = result[k].Get("billdate").Str
			userDetail.Name = result[k].Get("memname").Str
			userDetail.MethodPayment = result[k].Get("payway").Str
			userDetail.ShareProportion = result[k].Get("share_rate").Float()
			userDetail.Detail = result[k].Get("comboname").Str

			userCourse.Tickets = append(userCourse.Tickets, userDetail)

		}
	}
	for k := range userCourse.Tickets {
		userCourse.Total += userCourse.Tickets[k].ActualAmount
	}
	return
}

func GetUserResults(token string, shopId string, startUserId string, endUserId string, dates ...string) gjson.Result {
	var startTime, endTime string

	if len(dates) <= 1 {
		year := time.Now().Format("2006")
		month := time.Now().Format("01")
		day := time.Now().Format("02")
		startTime = fmt.Sprintf("%v%v01", year, month)
		endTime = fmt.Sprintf("%v%v%v", year, month, day)
	}

	if len(dates) == 2 {
		startTime = dates[0]
		endTime = dates[1]
	}

	client := ClientOption{
		Url:    "https://api.bokao2o.com/s3nos_report/person/v2/empPerformStats?v=1&sign=UERLUUcjMDAy",
		Params: nil,
		Headers: map[string]string{
			"access_token": token,
			"accesstoken":  shopId,
			"device_id":    "s3backend",
			"deviceid":     "s3backend",
			"referer":      "https://s3.boka.vc/home",
		},
		Body: map[string]interface{}{
			"compid":      "002",
			"compName":    "孔雀宫-迎春路",
			"fromdate":    startTime,
			"todate":      endTime,
			"fromempl":    startUserId,
			"toempl":      endUserId,
			"inc_card":    1,
			"inc_service": 1,
			"inc_goods":   1,
			"return_type": 1,
			"paymode":     "",
			"cardtype":    "",
			"recalculate": true,
			"type":        "2",
			"userId":      "ADMIN",
		},
	}
	res := client.Post()
	return client.ToJson(res)
}
