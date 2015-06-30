package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strconv"
	"time"
)

//将时间格式转为yyyymmdd,例如20150403
func ParseDateDay(date time.Time) string {

	year := strconv.Itoa(date.Year())
	month := strconv.Itoa(int(date.Month()))
	day := strconv.Itoa(date.Day())

	if int(date.Month()) < 10 {
		month = "0" + month
	}
	if date.Day() < 10 {
		day = "0" + day
	}

	return year + month + day
}

//将时间格式转为yyyymm,例如201504
func ParseDateMonth(date time.Time) string {

	year := strconv.Itoa(date.Year())
	month := strconv.Itoa(int(date.Month()))

	if int(date.Month()) < 10 {
		month = "0" + month
	}

	return year + month
}

//将时间格式转为hhmmss,例如120110
func ParseTime(date time.Time) string {

	hour := strconv.Itoa(date.Hour())
	minute := strconv.Itoa(date.Minute())
	second := strconv.Itoa(date.Second())

	if date.Hour() < 10 {
		hour = "0" + hour
	}
	if date.Minute() < 10 {
		minute = "0" + minute
	}
	if date.Second() < 10 {
		second = "0" + second
	}

	return hour + minute + second
}

//将时间格式转为yyyymmddhhmmss,例如20150408173301
func ParseDateTime(date time.Time) string {
	return ParseDateDay(date) + ParseTime(date)
}

//对字符转进行md5加密
func GeneMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成guid
func GeneGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GeneMd5(base64.URLEncoding.EncodeToString(b))
}
