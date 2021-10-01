package home

import (
	"go-blog/models/admin"
	"unsafe"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Log("index")

	// 推荐
	o := orm.NewOrm()
	var list []*admin.Article
	o.QueryTable(new(admin.Article)).Filter("status", 1).Filter("recommend", 1).Filter("User__Name__isnull", false).Filter("Category__Name__isnull", false).OrderBy("-id").RelatedSel().All(&list, "id", "title")
	c.Data["Recommend"] = list

	c.Data["index"] = "首页"

	var topic []*admin.Topic
	o.QueryTable(new(admin.Topic)).Filter("status", 1).OrderBy("-join").Limit(10).All(&topic)
	c.Data["Topic"] = topic

	if beego.AppConfig.String("view") == "nihongdengxia" {
		((*ArticleController)(unsafe.Pointer(c))).List()
	}

	c.TplName = "home/" + c.Template + "/index.html"
}
