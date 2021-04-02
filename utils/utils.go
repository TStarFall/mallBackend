package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

//页面显示限制
func Page(Limit, Page int)(limit, offset int) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

//排序
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}

//定义时间戳常量
const TimeLayOut = "2006-01-02 15:04:05"

var (
	Local = time.FixedZone("CST", 8*3600)
)

//获取当前时间
func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayOut)
	return now
}

//时间格式化
func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayOut, s, time.Local)
	if err != nil {
		panic(err)
	}
	return result.In(Local).Format(TimeLayOut)
}

//MD5加密串
func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x",w.Sum(nil))
	return md5Str
}








