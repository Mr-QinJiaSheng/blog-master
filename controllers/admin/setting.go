package admin

import (
	"go-blog/models/admin"
	"go-blog/utils"

	"github.com/astaxie/beego/orm"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Add() {

	o := orm.NewOrm()
	var setting []*admin.Setting
	o.QueryTable(new(admin.Setting)).All(&setting)

	for _, v := range setting {
		c.Data[v.Name] = v.Value
	}

	path := utils.GetViewPaths()
	c.Data["View"] = path
	c.TplName = "admin/setting.html"
}

func (c *SettingController) Save() {

	response := make(map[string]interface{})

	title := c.GetString("title")
	keyword := c.GetString("keyword")
	description := c.GetString("description")
	//name := c.GetString("name")
	tag := c.GetString("tag")
	//remark := c.GetString("remark")
	//remark_markdown_doc := c.GetString("remark-markdown-doc")
	//remark_html_code := c.GetString("remark-markdown-doc")
	image := c.GetString("image")

	template := c.GetString("template")
	limit := c.GetString("limit")

	o := orm.NewOrm()
	err := o.Begin()

	//_, err = o.Delete(&admin.Setting{Name: "name"})
	_, err = o.Delete(&admin.Setting{Name: "title"})
	_, err = o.Delete(&admin.Setting{Name: "tag"})
	_, err = o.Delete(&admin.Setting{Name: "template"})
	//_, err = o.Delete(&admin.Setting{Name: "remark"})
	_, err = o.Delete(&admin.Setting{Name: "image"})
	_, err = o.Delete(&admin.Setting{Name: "keyword"})
	_, err = o.Delete(&admin.Setting{Name: "description"})
	//_, err = o.Delete(&admin.Setting{Name: "remark_markdown_doc"})
	//_, err = o.Delete(&admin.Setting{Name: "remark_html_code"})
	_, err = o.Delete(&admin.Setting{Name: "limit"})

	settings := []admin.Setting{
		{Name: "title", Value: title},
		{Name: "template", Value: template},
		//{Name: "name", Value: name},
		{Name: "limit", Value: limit},
		{Name: "tag", Value: tag},
		//{Name: "remark_markdown_doc", Value: remark_markdown_doc},
		//{Name: "remark_html_code", Value: remark_html_code},
		{Name: "image", Value: image},
		{Name: "keyword", Value: keyword},
		{Name: "description", Value: description},
	}

	num, err := o.InsertMulti(7, settings)

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	} else {
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *SettingController) Notice() {

	o := orm.NewOrm()
	var setting admin.Setting
	o.QueryTable(new(admin.Setting)).Filter("name", "notice").One(&setting)
	c.Data["Notice"] = setting.Value
	c.TplName = "admin/notice.html"
}

func (c *SettingController) NoticeSave() {

	response := make(map[string]interface{})

	notice := c.GetString("notice")

	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(&admin.Setting{Name: "notice"})

	num, err := o.Insert(&admin.Setting{Name: "notice", Value: notice})

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	} else {
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *SettingController) About() {

	o := orm.NewOrm()

	var setting []*admin.Setting
	o.QueryTable(new(admin.Setting)).All(&setting)

	for _, v := range setting {
		c.Data[v.Name] = v.Value
	}
	c.TplName = "admin/about.html"
}

func (c *SettingController) AboutSave() {

	response := make(map[string]interface{})

	//about := c.GetString("about")
	about_markdown_doc := c.GetString("about-markdown-doc")
	about_html_code := c.GetString("about-html-code")

	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(&admin.Setting{Name: "about_markdown_doc"})
	_, err = o.Delete(&admin.Setting{Name: "about_html_code"})

	settings := []admin.Setting{
		{Name: "about_markdown_doc", Value: about_markdown_doc},
		{Name: "about_html_code", Value: about_html_code},
	}

	num, err := o.InsertMulti(2, settings)

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	} else {
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
