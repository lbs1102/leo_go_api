package common

import (
	"fmt"
	"os"
	"runtime"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var date_arr map[string]string
var log *logrus.Logger

func LogTrace(msg string) {
	set_log_dir("Trace")
	log.Tracef("%s [%s] - %s", GetDateTime(), GetActionName(), msg)
}

func LogDebug(msg string) {
	set_log_dir("Debug")
	log.Debugf("%s [%s] - %s", GetDateTime(), GetActionName(), msg)
}

func LogInfo(msg string) {
	set_log_dir("Info")
	log.Infof("%s [%s] - %s", GetDateTime(), GetActionName(), msg)
}

func LogWarn(msg string) {
	set_log_dir("Warn")

	log.Warnf("%s [%s] - %s", GetDateTime(), GetActionName(), msg)
}

func LogError(msg string) {
	set_log_dir("Error")
	log.Errorf("%s [%s] - %s", GetDateTime(), GetActionName(), msg)
}

func LogPanic(msg string) {
	set_log_dir("Panic")
	log.Panicf("%s [%s] - %s", GetDateTime(), GetActionName(), msg)
}

func GetActionName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func set_log_dir(active string) {
	/*leo 建立各類別的log路徑
	* 目前會再根目錄等下建立log資料夾，在依照各類別建立相對應的log資料夾
	 */
	date_arr = GetNowDateArray()
	dir, _ := os.Getwd()
	log_name := fmt.Sprintf("%s_%s.log", date_arr["hour"], date_arr["min"])
	date_dir := fmt.Sprintf("%s/%s/%s", date_arr["year"], date_arr["month"], date_arr["day"])
	dir = fmt.Sprintf("%s/%s/%s/%s", dir, OutputDir, active, date_dir)
	os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	dir_file := fmt.Sprintf("%s/%s", dir, log_name)
	output, _ := os.OpenFile(dir_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(output)
}

func Initlogger() *logrus.Logger {
	log = &logrus.Logger{
		Out: &lumberjack.Logger{
			MaxSize:    5, // megabytes
			MaxBackups: 90,
			MaxAge:     90,    //days
			Compress:   false, // disabled by default
		},
		Level: logrus.TraceLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors:    true,
			DisableQuote:     true,
			DisableTimestamp: true,
		},
	}
	return log
}
