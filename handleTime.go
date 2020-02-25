package handleTime

import "time"

//time.Time 类型的时间格式转化为相应的时间格式(传入参数为time.now)

//type location time.Local
var location=time.Local
const  timeLayout = "2006-01-02 15:04:05"
const timeDate  = "2006-01-02"
//设置相应的时区
func GetLocation(local string){
	location,_ = time.LoadLocation(local)
}

//时间格式转化为相应的日期格式,如传入time.now()
func TimeToYYYYMMDDHHMMSS(t time.Time)string{
	//loc,_:=time.LoadLocation("Asia/Beijing")
	timeNow:=t.Format(timeLayout)
	return timeNow
}

func TimeToYYYYMMDD(t time.Time) string{
	timeNow:=t.Format(timeDate)
	return timeNow
}

//时间戳转化为时间
func TimeStampToTime(t int64)string{
	datetime := time.Unix(t, 0).Format(timeLayout)
	return datetime
}

//string格式转化为时间戳格式
func StringToYYYYMMDD(str string)(time.Time,error){
	return time.Parse(timeDate,str)
}

func StringToYYYYMMDDHHMMSS(str string)(time.Time,error){
	return time.Parse(timeLayout,str)
}

//时间差
func TimeSub(t1 int64,t2 int64)string{
	time1:= time.Unix(t1,0)
	time2:=time.Unix(t2,0)
	subTime:= time1.Sub(time2)

}
