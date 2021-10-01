package admin

import (
	"encoding/json"
	"fmt"
	"go-blog/models/admin"
	"go-blog/utils"
	"go-blog/utils/article"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ArticleResourcesController struct {
	BaseController
}

func (c *ArticleResourcesController) GetArticle() {

	t1 := time.Now()
	response := make(map[string]interface{})
	urlPath := c.GetString("url")
	//aid, err := c.GetInt("aid")
	if urlPath == "" {
		response["msg"] = "抓取失败！"
		response["code"] = 500
		response["err"] = "url must no be null"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	u, err := url.Parse(urlPath)
	if err != nil {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}
	host := u.Host
	ho := strings.Split(host, ":")

	//var title,html,md string
	var data map[string]interface{}
	var md, title, cover, html string

	//fmt.Println(ho[0])
	switch ho[0] {
	case "gocn.vip":
		gocn := article.Gocn{}
		data = gocn.Get(urlPath)
		title = data["title"].(string)
		html = data["content"].(string)
		md = utils.Html2md(data["content"].(string))
	case "book.douban.com":
		douban := article.Douban{}
		data = douban.Get(urlPath)
		subject := data["subject"].(map[string]string)
		review := data["list"].([]map[string]string)
		cover = subject["src"]
		subjectJson, _ := json.Marshal(subject)
		subjectString := string(subjectJson)
		//fmt.Println(subject)

		detail, err := douban.GetDetail(urlPath)
		//fmt.Println(detail)
		if err != nil {
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = "获取文章失败!"
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
		html = detail["content"].(string)
		md = utils.Html2md(detail["content"].(string))

		o := orm.NewOrm()
		art := admin.Article{
			Cover:    cover,
			Review:   len(review),
			Title:    detail["title"].(string),
			Remark:   detail["desc"].(string),
			Desc:     md,
			Html:     html,
			Category: &admin.Category{Id: 1},
			User:     &admin.User{Id: 1},
			Status:   1,
			Other:    subjectString,
			Tag:      detail["title"].(string),
		}

		id, err := o.Insert(&art)
		if err != nil {
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = err.Error()
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
		aid := int(id)

		if aid == 0 {
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = "文章ID不能为空!"
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}

		count := 0
		//var m []admin.Review
		worker, err := utils.NewIdWorker(1)
		for _, v := range review {

			uid, err := worker.GetNextId()
			if err != nil {
				uid, err = worker.GetNextId()
				if err != nil {
					uid = time.Now().Unix()
				}
			}
			uidstr := strconv.FormatInt(uid, 10)
			customerId, _ := admin.AddCustomer(&admin.Customer{
				Uid:      uidstr,
				Username: v["username"],
				Nickname: v["username"],
				Image:    v["avatar"],
			})
			rev := admin.Review{}
			rev.ArticleId = aid
			rev.Review = v["content"]
			rev.Name = v["username"]
			rev.Status = 1
			rev.Customer = &admin.Customer{Id: int(customerId)}
			rating, _ := strconv.Atoi(v["rating"])
			//fmt.Println("星星：", rating)
			rev.Star = rating
			like, _ := strconv.Atoi(v["like"])
			rev.Like = like
			loc, _ := time.LoadLocation("Local")

			//fmt.Println("时间：", v["commenttime"])
			the_time, _ := time.ParseInLocation("2006-01-02 15:04:05", v["commenttime"], loc)
			rev.Created = the_time

			// s["content"] = e.ChildText(".comment-content")
			// s["username"] = e.ChildText("h3 .comment-info a")
			// s["avatar"] = e.ChildAttr(".avatar a img", "src")
			// s["commenttime"] = e.ChildText("h3 .comment-info .comment-time")
			// s["rating"] = e.ChildAttr(".user-stars", "class")
			// type Review struct {
			// 	Id        int
			// 	Name      string
			// 	Review    string    `orm:"size(500)"`
			// 	Reply     string    `orm:"size(500)"`
			// 	Site      string    `orm:"size(500)"`
			// 	Created   time.Time `orm:"auto_now_add;type(datetime)"`
			// 	Updated   time.Time `orm:"auto_now;type(datetime)"`
			// 	Status    int       `orm:"default(1)"`
			// 	ArticleId int
			// 	Customer  *Customer `orm:"rel(fk)"`
			// 	Like      int
			// 	Star      int
			// }
			if _, err := o.Insert(&rev); err != nil {
				fmt.Println(err.Error())
				continue
			}
			//m = append(m, rev)
			count++
		}

		//fmt.Println(m)
		//fmt.Println(aid)
		//err = admin.AddReview(m, art.Id)

		response["msg"] = "总共评论 " + strconv.Itoa(len(review)) + "条，爬取成功 " + strconv.Itoa(count) + "条！"
		response["code"] = 200
		response["id"] = aid
		response["title"] = art.Title
		response["elapsed"] = fmt.Sprintf("%s", time.Since(t1))
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()

	default:
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = "未知源!"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	o := orm.NewOrm()
	article := admin.Article{
		Title:    title,
		Tag:      "",
		Desc:     md,
		Html:     html,
		Remark:   "",
		Status:   1,
		User:     &admin.User{1, "", "", "", time.Now(), 0},
		Category: &admin.Category{1, "", 0, 0, 0},
		Cover:    cover,
	}

	if id, err := o.Insert(&article); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		response["id"] = id
		response["title"] = title
		response["elapsed"] = fmt.Sprintf("%s", time.Since(t1))
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *ArticleResourcesController) GetCron() {
	c.TplName = "admin/article-auto-add.html"
}
