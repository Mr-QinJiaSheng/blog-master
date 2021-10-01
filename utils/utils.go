package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func GetViewPaths() (name []string) {
	dir, err := ioutil.ReadDir("views/home")
	if err != nil {
		fmt.Println(err)
		return name
	}

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			name = append(name, fi.Name())

		}
	}

	return name
}

func CreateUsername() string {

	// 形容词
	jsonFile, err := os.Open("utils/basedata/adjective.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}
	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var adjective []string
	json.Unmarshal([]byte(byteValue), &adjective)

	// 名称
	jsonFile, err = os.Open("utils/basedata/noun.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}
	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ = ioutil.ReadAll(jsonFile)

	var nonu []string
	json.Unmarshal([]byte(byteValue), &nonu)

	rand.Seed(time.Now().Unix())
	return adjective[rand.Intn(len(adjective))] + nonu[rand.Intn(len(nonu))]
}

func Subbbs(source string, l int) string {
	var r = []rune(source)
	length := len(r)
	if l >= length {
		return source
	}

	return string(r[0:(l-4)]) + `...<a href="#">(展开)</a>`
}

func GetTopic(s string) (topic, content string) {
	r := []rune(s)
	if string(r[0]) != "#" {
		return "", s
	}

	l := 0
	for _, v := range r {

		topic += string(v)

		if string(v) == "#" {
			l += 1
		}
		if l == 2 {
			break
		}

	}
	return topic, string(r[len([]rune(topic)):])
}

// 下载图片
func DownImage(imgUrl string) (filename string, err error) {
	//fmt.Println(beego.WorkPath)
	appPath, _ := os.Getwd()
	fileBaseName := path.Base(imgUrl)
	imgPath := filepath.Join("/static/uploads", time.Now().Format("2006-01"))
	filename = filepath.Join(imgPath, fileBaseName)
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if err = CheckDir(appPath + imgPath); err != nil {
		return
	}
	f, err := os.OpenFile(appPath+filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return
	} else {
		_, err = f.Write(b)
		if err != nil {
			return
		}
	}
	return

	// err = ioutil.WriteFile(imgPath, b, os.ModePerm)
	// return
}

func CheckDir(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			return err
		}
	}

	// check again
	_, err := os.Stat(path)
	return err
}
