package requests

import (
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)

type UsersCourse struct {
	Course map[string]UserCourse `json:"course"`
}

type UserCourse struct {
	User    string       `json:"user"`
	Total   int64        `json:"total"`
	Consume int64        `json:"consume"`
	Case    int64        `json:"case"`
	Tickets []UserDetail `json:"tickets"`
}
type UserDetail struct {
	Date            string  `json:"date"`
	Name            string  `json:"name"`
	Amount          int64   `json:"amount"`
	ActualAmount    int64   `json:"actualAmount"`
	MethodPayment   string  `json:"methodPayment"`
	ShareProportion float64 `json:"shareProportion"`
	Detail          string  `json:"detail"`
}

func (user *UsersCourse) GetCourses(data gjson.Result) {

	if data.Get("code").Int() != 200 {
		return
	}

	result := data.Get("result").Array()

	user.Course = make(map[string]UserCourse)

	for _, k := range result {
		userId, course := user.GetCourse(k)
		if len(userId) == 3 {
			user.Course[userId] = course
		}

	}
	return
}

func (user *UsersCourse) GetCourse(data gjson.Result) (personId string, userCourse UserCourse) {

	for _, result := range data.Array() {
		actionID := result.Get("action_id").Int()
		// 获取员工工号
		if personId == "" && result.Get("person_id").Int() != 0 {
			personId = result.Get("person_id").String()
			userCourse.User = personId
		}

		if (actionID == 0 || actionID == 1) && result.Get("person_id").Int() != 0 {
			//a := result[k].Get("amt").Num
			var userDetail UserDetail

			userDetail.Amount = result.Get("amt").Int()
			userDetail.ActualAmount = result.Get("amt3").Int()
			userDetail.Date = result.Get("billdate").Str
			userDetail.Name = result.Get("memname").Str
			userDetail.MethodPayment = result.Get("payway").Str
			userDetail.ShareProportion = result.Get("share_rate").Float()

			linshibianliang := result.Get("comboname").Str
			if linshibianliang == "" {
				linshibianliang = "充值或卖卡"

			} else {
				linshibianliang = "课程: " + linshibianliang
			}
			userDetail.Detail = linshibianliang

			userCourse.Tickets = append(userCourse.Tickets, userDetail)

		}

		// 课程消耗
		if (actionID == 3 || actionID == 6) && result.Get("person_id").Int() != 0 {
			if result.Get("payway").Str == "疗程账户" {
				userCourse.Consume += result.Get("amt3").Int()
			}
		}
	}

	for k := range userCourse.Tickets {
		if userCourse.Tickets[k].Detail != "充值或卖卡" {
			userCourse.Total += userCourse.Tickets[k].ActualAmount
		} else {
			userCourse.Case += userCourse.Tickets[k].ActualAmount
		}

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
