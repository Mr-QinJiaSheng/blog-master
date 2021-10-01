package utils

import (
	"bytes"
	"math"
	"strconv"
	"strings"
	"time"
)

func IndexForOne(i int, p, limit int64) int64 {
	s := strconv.Itoa(i)
	index, _ := strconv.ParseInt(s, 10, 64)
	return (p-1)*limit + index + 1
}

func IndexAddOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index + 1
}

func IndexDecrOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index - 1
}

func StringReplace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func StringToTime(date interface{}) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	ret, _ := time.ParseInLocation(timeLayout, date.(string), loc)
	return ret
}

func TimeStampToTime(timeStamp int32) time.Time {
	return time.Unix(int64(timeStamp), 0)
}

func TemlpateTime(t time.Time) string {

	atime := t.Unix()
	var byTime = []int64{24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"天前", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 30 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]

		if i == 0 && temp > 1 {
			res = t.Format("2006-01-02 15:04")
			break
		}
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = mergeString(tempStr, unit[i]) //此处调用了一个我自己封装的字符串拼接的函数（你也可以自己实现）
		}
		break //我想要的形式是精确到最大单位，即："2天前"这种形式，如果想要"2天12小时36分钟48秒前"这种形式，把此处break去掉，然后把字符串拼接调整下即可（别问我怎么调整，这如果都不会我也是无语）
	}
	return res
}

/**
* @des 拼接字符串
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func mergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}
