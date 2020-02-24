package handleTime

import "time"

//time.Time 类型的时间格式转化为相应的时间格式(传入参数为time.now)

//type location time.Local
var location=time.Local

//设置相应的时区
func GetLocation(local string){
	location,_ = time.LoadLocation(local)
}

//时间格式转化为相应的日期格式,如传入time.now()
func TimeToYYYYMMDDHHMMSS(t time.Time)string{
	//loc,_:=time.LoadLocation("Asia/Beijing")
	timeNow:=t.Format("2006-01-02 15:04:05")
	return timeNow
}

func TimeToYYYYMMDD(t time.Time) string{
	timeNow:=t.Format("2006-01-02")
	return timeNow
}

//时间戳转化为时间
func TimeStampToTime(t int64){

}

//string格式转化为时间戳格式
func GetTimeStamp(str string)int64{
	return 1
}

