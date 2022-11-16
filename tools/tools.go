package tools

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

func TimeNow10() int64 {
	now := time.Now().UTC().Unix()
	return now
}

func TimeNow13() int64 {
	now := time.Now().UTC().UnixMilli()
	return now
}

func Timestamp10ToDatetime(timestamp int64) string {
	nt := time.Unix(timestamp, 0)
	const base_format = "2006-01-02 15:04:05"
	return nt.Format(base_format)
}
func Timestamp13ToDatetime(timestamp int64) string {
	nt := time.Unix(timestamp/1000, 0)
	const base_format = "2006-01-02 15:04:05"
	return nt.Format(base_format)
}

func DatetimeToTimestamp(datetime string) int64 {
	const base_format = "2006-01-02 15:04:05"
	temp, _ := time.Parse(base_format, datetime)
	return temp.Unix()
}

func CheckMysqlErr(err error) {
	if err != nil {
		log.Println("mysql error", err.Error())
	}
}

// # 格式轉換

func StringToFloat64(info string) float64 {
	tmp, _ := strconv.ParseFloat(info, 64)
	return tmp
}

func StringToInt(info string) int {
	tmp, _ := strconv.Atoi(info)
	return tmp
}

func StringToInt64(info string) int64 {
	tmp, _ := strconv.ParseInt(info, 10, 64)
	return tmp
}

func IntToString(info int) string {
	tmp := strconv.Itoa(info)
	return tmp
}

func Int64ToString(info int64) string {
	tmp := strconv.FormatInt(info, 10)
	return tmp
}

func Float64ToString(info float64) string {
	tmp := fmt.Sprintf("%f", info)
	return tmp
}

func StringToTime(info string) time.Time {
	tmp := time.Unix(StringToInt64(info[:10]), 0)
	return tmp
}

// 密碼md5加密
func PasswdSetMd5(passwd string) string {
	w := md5.New()
	io.WriteString(w, passwd)               //将str写入到w中
	md5str := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5str
}

func CheckErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
