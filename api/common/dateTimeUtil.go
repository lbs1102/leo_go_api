package common

import (
	"fmt"
	"time"
)

func GetTime() int64 {
	return time.Now().Unix()
}

func GetDateTime() string {
	var ts int64 = GetTime()
	var tm time.Time = time.Unix(ts, 0)
	return tm.Format("2006-01-02 15:04:05")
}
func GetNowDateArray() map[string]string {
	res := make(map[string]string)
	res["year"] = fmt.Sprintf("%v", time.Now().Year())
	res["month"] = fmt.Sprintf("%v", int(time.Now().Month()))
	res["day"] = fmt.Sprintf("%v", time.Now().Day())
	res["hour"] = fmt.Sprintf("%v", time.Now().Hour())
	res["min"] = fmt.Sprintf("%v", time.Now().Minute())
	res["sec"] = fmt.Sprintf("%v", time.Now().Second())
	return res
}
