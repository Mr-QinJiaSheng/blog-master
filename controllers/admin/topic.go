package admin

import (
	"go-blog/models/admin"
	"go-blog/utils"

	"github.com/astaxie/beego/orm"
)

type TopicController struct {
	BaseController
}

func (c *TopicController) List() {
	o := orm.NewOrm()

	limit := int64(15)
	page, _ := c.GetInt64("page", 1) // 页数
	offset := (page - 1) * limit     // 偏移量

	var topic []*admin.Topic
	qs := o.QueryTable(new(admin.Topic))
	qs.Limit(limit).Offset(offset).All(&topic)

	for _, v := range topic {
		var category admin.Category
		if v.CategoryId != 0 {
			o.QueryTable(new(admin.Category)).Filter("id", v.CategoryId).One(&category)
		}
		v.Category = &category
	}
	// 统计
	count, _ := qs.Count()

	c.Data["Data"] = topic
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.TplName = "admin/topic.html"
}

func (c *TopicController) Put() {
	id, err := c.GetInt("id", 0)

	if id == 0 {
		c.Abort("404")
	}

	// 基础数据
	o := orm.NewOrm()
	var topic admin.Topic
	qs := o.QueryTable(new(admin.Topic))
	err = qs.Filter("id", id).One(&topic)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = topic

	// 分类数据
	o = orm.NewOrm()
	category := new(admin.Category)
	var categorys []*admin.Category
	qs = o.QueryTable(category)
	qs = qs.Filter("status", 1)
	qs.All(&categorys)

	c.Data["Category"] = categorys
	/*c.Data["json"]= &articles
	c.ServeJSON()
	c.StopRun()*/

	c.TplName = "admin/topic-edit.html"

}

func (c *TopicController) Update() {
	id, _ := c.GetInt("id", 0)
	category_id, _ := c.GetInt("category_id", 0)
	status, _ := c.GetInt("status", 0)

	/*c.Data["json"] = c.Input()
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	o := orm.NewOrm()

	topic := admin.Topic{Id: id}
	if o.Read(&topic) == nil {
		topic.CategoryId = category_id
		topic.Status = status

		if _, err := o.Update(&topic); err == nil {
			response["msg"] = "修改成功！"
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "修改失败！"
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "修改失败！"
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
