package article

import (
	"encoding/json"
	"fmt"
	"go-blog/utils"
	"strings"

	"github.com/gocolly/colly"
)

type Douban struct {
}

func (douban Douban) Get(url string) map[string]interface{} {

	url = url + "comments/"
	var sli []map[string]string
	subject := make(map[string]string)
	c := colly.NewCollector()
	c.OnHTML("#comments .comment-list .comment-item", func(e *colly.HTMLElement) {
		s := make(map[string]string)
		s["content"] = e.ChildText(".comment-content")
		s["username"] = e.ChildText("h3 .comment-info a")
		s["avatar"] = e.ChildAttr(".avatar a img", "src")
		s["commenttime"] = e.ChildText("h3 .comment-info .comment-time")
		s["rating"] = e.ChildAttr(".user-stars", "class")
		s["like"] = e.ChildText("h3 .comment-vote .vote-count")
		sli = append(sli, s)
		//fmt.Println(s)
	})

	c.OnHTML(".aside .subject-info", func(h *colly.HTMLElement) {
		imgsrc := h.ChildAttr("a img", "src")

		img, err := utils.DownImage(imgsrc)
		if err != nil {
			fmt.Println(err.Error())
		}
		subject["src"] = img
		subject["title"] = h.ChildAttr("a img", "title")
		subjectInfo := h.Text
		// 去除空格
		subjectInfo = strings.Replace(subjectInfo, " ", "", -1)
		// 去除换行符
		//subjectInfo = strings.Replace(subjectInfo, "\n", "", -1)
		//fmt.Println(src)
		//fmt.Println(subjectInfo)

		str := strings.Split(subjectInfo, "\n")
		for _, v := range str {
			if v != "" {
				subject["subjectInfo"] += v + "\n"
				info := strings.Split(v, ":")

				if info[0] == "作者" {
					subject["author"] = info[1]
				}

				if info[0] == "书名" {
					subject["bookname"] = info[1]
				}
			}
		}

		//fmt.Println(subject)
	})

	//subject["subjectInfo"] = strings.TrimSpace(subject["subjectInfo"])
	// On every a element which has href attribute call callback
	// c.OnHTML(".col-md-9 .topic-detail .card-body", func(e *colly.HTMLElement) {
	// 	//fmt.Println(e.DOM.Children().Filter(".card").Remove())
	// 	e.DOM.Children().Filter(".toc-container").Remove()
	// 	//d.ChildrenFiltered()
	// 	e.DOM.Children().Filter(".card").Remove()
	// 	s, _ = e.DOM.Html()
	// })
	// Start scraping on https://hackerspaces.org
	c.Visit(url)
	ret := make(map[string]interface{})
	ret["subject"] = subject
	ret["list"] = sli
	//fmt.Println(ret)
	return ret
}

func (douban Douban) GetDetail(url string) (map[string]interface{}, error) {
	url = url + "reviews/?sort=hotest"
	c := colly.NewCollector()
	ret := make(map[string]interface{})
	var id string
	c.OnHTML(".review-list>div:first-child", func(e *colly.HTMLElement) {

		ret["desc"] = e.ChildText(".main-bd .review-short .short-content")
		ret["title"] = strings.TrimSpace(e.ChildText(".main-bd h2"))
		id = e.Attr("data-cid")

	})

	c.Visit(url)

	resp := getArticle("https://book.douban.com/j/review/" + id + "/full")

	ret["content"] = strings.Replace(resp.Html, "\n", "<br>", -1)

	//fmt.Println(ret)
	return ret, nil

}

type BaseJsonBean struct {
	Html string `json:"html"`
}

func getArticle(url string) BaseJsonBean {

	c := colly.NewCollector()
	var result BaseJsonBean
	c.OnResponse(func(r *colly.Response) {
		json.Unmarshal(r.Body, &result)
	})

	c.Visit(url)

	return result

}
