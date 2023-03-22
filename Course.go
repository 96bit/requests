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

	if data.Get("code").Int() != 200 {
		return
	}

	result := data.Get("result").Array()

	user.Course = make(map[string]UserCourse)

	for _, k := range result {
		userId, course := user.GetCourse(k)
		if len(userId) == 3 {
			fmt.Println(course)
			user.Course[userId] = course
		}

	}
	return
}

func (user *UsersCourse) GetCourse(data gjson.Result) (personId string, userCourse UserCourse) {

	for _, result := range data.Array() {
		if result.Get("action_id").Int() == 0 && result.Get("person_id").Int() != 0 {
			if personId == "" {
				personId = result.Get("person_id").String()
				userCourse.User = personId
			}
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
