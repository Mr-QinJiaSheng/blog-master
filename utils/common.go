package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go-blog/models/admin"
	"reflect"
	"strconv"
	"time"
)

// @Title 生成密码
// @Description create AccountAccount
// @Param	body		body 	models.AccountAccount	true		"body for AccountAccount content"
// @Success 201 {int} models.AccountAccount
// @Failure 403 body is empty
func PasswordMD5(passwd, salt string) string {
	h := md5.New()
	// 后面增加一个无意义字符串
	h.Write([]byte(passwd + salt + "@.YnO-"))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}

// ToString 类型转换，获得string
func ToString(v interface{}) (re string) {
	re = v.(string)
	return
}

// StringsJoin 字符串拼接
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}

// ToInt64 类型转换，获得int64
func ToInt64(v interface{}) (re int64, err error) {
	switch v.(type) {
	case string:
		re, err = strconv.ParseInt(v.(string), 10, 64)
	case float64:
		re = int64(v.(float64))
	case float32:
		re = int64(v.(float32))
	case int64:
		re = v.(int64)
	case int32:
		re = v.(int64)
	default:
		err = errors.New("不能转换")
	}
	return
}

// ToSlice 转换为数组
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

type CateTree struct {
	Id    int
	Name  string
	Pid   int
	Sort  int
	Level int
	Son   []CateTree
}

//递归实现(返回树状结果得数据)
func CategoryTree(allCate []*admin.Category, pid int, level int) []CateTree {
	var arr []CateTree
	for _, v := range allCate {
		if pid == v.Pid {
			ctree := CateTree{}
			ctree.Id = v.Id
			ctree.Pid = v.Pid
			ctree.Name = v.Name
			ctree.Sort = v.Sort
			ctree.Level = level
			sonCate := CategoryTree(allCate, v.Id, level+1)
			ctree.Son = sonCate
			arr = append(arr, ctree)
		}
	}
	return arr
}

func CategoryTreeR(allCate []*admin.Category, pid int, level int) []CateTree {
	var arr []CateTree
	for _, v := range allCate {
		if pid == v.Pid {
			ctree := CateTree{}
			ctree.Id = v.Id
			ctree.Pid = v.Pid
			ctree.Name = v.Name
			ctree.Sort = v.Sort
			ctree.Level = level
			arr = append(arr, ctree)
			sonCate := CategoryTreeR(allCate, v.Id, level+1)
			arr = append(arr, sonCate...)
			//ctree.Son = sonCate
			//arr = append(arr, ctree)
		}
	}
	return arr
}

func SubString(str string, len int) string {
	return string([]rune(str)[:len])
}

type MenuTree struct {
	Id     int
	Title  string
	Pid    int
	Sort   int
	Level  int
	Url    string
	Target string
	Son    []MenuTree
}

func MenuTreeR(allCate []interface{}, pid int, level int) []MenuTree {
	var arr []MenuTree
	for _, v := range allCate {
		v1 := v.(admin.Menu)
		if pid == v1.Pid {
			ctree := MenuTree{}
			ctree.Id = v1.Id
			ctree.Pid = v1.Pid
			ctree.Title = v1.Title
			ctree.Sort = v1.Sort
			ctree.Level = level
			ctree.Url = v1.Url
			ctree.Target = v1.Target
			arr = append(arr, ctree)
			sonCate := MenuTreeR(allCate, v1.Id, level+1)
			arr = append(arr, sonCate...)
		}
	}
	return arr
}

func MenuData(allCate []interface{}, pid int, level int) []MenuTree {
	var arr []MenuTree
	for _, v := range allCate {
		v1 := v.(admin.Menu)
		if pid == v1.Pid {
			ctree := MenuTree{}
			ctree.Id = v1.Id
			ctree.Pid = v1.Pid
			ctree.Title = v1.Title
			ctree.Sort = v1.Sort
			ctree.Level = level
			ctree.Url = v1.Url
			ctree.Target = v1.Target
			sonCate := MenuData(allCate, v1.Id, level+1)
			ctree.Son = sonCate
			arr = append(arr, ctree)
		}
	}
	return arr
}

type ReviewTree struct {
	Id         int
	Username   string
	Image      string
	Content    string
	ReplyID    int
	ReplyName  string
	ReplyImage string
	Created    time.Time
	Level      int
	Son        []ReviewTree
}

// [
//     {
//         "Id":1,
//         "Username":"我是甜美的西红柿",
//         "Content":"ssss",
//         "ReplyID":0,
//         "ReplyName":"",
//         "Creared":"2020-12-05T00:11:26+08:00",
//         "Level":0,
//         "Son":[
//             {
//                 "Id":2,
//                 "Username":"3乐",
//                 "Content":"wefsf",
//                 "ReplyID":1,
//                 "ReplyName":"我是甜美的西红柿",
//                 "Creared":"2020-12-05T00:11:50+08:00",
//                 "Level":1,
//                 "Son":null
//             },
//             {
//                 "Id":3,
//                 "Username":"小丑向月亮生气",
//                 "Content":"qweqeqe",
//                 "ReplyID":2,
//                 "ReplyName":"3乐",
//                 "Creared":"2020-12-05T10:28:38+08:00",
//                 "Level":2,
//                 "Son":null
//             },
//             {
//                 "Id":4,
//                 "Username":"柴犬妹妹",
//                 "Content":"qewqeqeq",
//                 "ReplyID":1,
//                 "ReplyName":"我是甜美的西红柿",
//                 "Creared":"2020-12-05T10:28:48+08:00",
//                 "Level":1,
//                 "Son":null
//             }
//         ]
//     },
//     {
//         "Id":5,
//         "Username":"小莉啊",
//         "Content":"qewqeqeqqeqeqw",
//         "ReplyID":0,
//         "ReplyName":"",
//         "Creared":"2020-12-05T10:28:58+08:00",
//         "Level":0,
//         "Son":null
//     }
// ]
func ReviewTreeR(allCate []*admin.BbsReview, pid int, level int, self *admin.BbsReview) []ReviewTree {
	var arr []ReviewTree
	for _, v := range allCate {

		if pid == v.ReplyId {
			review := ReviewTree{}
			review.Id = v.Id
			review.ReplyID = v.ReplyId
			review.Created = v.Created
			review.Content = v.Content
			review.Level = level
			review.Username = v.Customer.Username
			review.Image = v.Customer.Image
			if pid != 0 {
				review.ReplyName = self.Customer.Username
				review.ReplyImage = self.Customer.Image
			}
			//arr = append(arr, review)
			if pid == 0 {
				sonReview := ReviewTreeR(allCate, v.Id, level+1, v)
				review.Son = sonReview
				arr = append(arr, review)
			} else {
				arr = append(arr, review)
				//review = ReviewTreeR(allCate, v.Id, level+1, v)
				arr = append(arr, ReviewTreeR(allCate, v.Id, level+1, v)...)
				//review.Son = append(review.Son, sonReview...)
			}
			//arr = append(arr, review)
			//arr = append(arr, review)
		}
	}
	return arr
}

// type AdTree struct {
// 	Gid        int
// 	Username  string
// 	Content   string
// 	ReplyID   int
// 	ReplyName string
// 	Created   time.Time
// 	Level     int
// 	Son       []ReviewTree
// }

// func AdTreeR(allCate []*admin.Ad, pid int, level int, self *admin.Ad) []AdTree {
// 	var arr []ReviewTree
// 	for _, v := range allCate {

// 		if pid == v.ReplyId {
// 			review := ReviewTree{}
// 			review.Id = v.Id
// 			review.ReplyID = v.ReplyId
// 			review.Created = v.Created
// 			review.Content = v.Content
// 			review.Level = level
// 			review.Username = v.Customer.Username
// 			if pid != 0 {
// 				review.ReplyName = self.Customer.Username
// 			}
// 			//arr = append(arr, review)
// 			if pid == 0 {
// 				sonReview := ReviewTreeR(allCate, v.Id, level+1, v)
// 				review.Son = sonReview
// 				arr = append(arr, review)
// 			} else {
// 				arr = append(arr, review)
// 				//review = ReviewTreeR(allCate, v.Id, level+1, v)
// 				arr = append(arr, ReviewTreeR(allCate, v.Id, level+1, v)...)
// 				//review.Son = append(review.Son, sonReview...)
// 			}
// 			//arr = append(arr, review)
// 			//arr = append(arr, review)
// 		}
// 	}
// 	return arr
// }
