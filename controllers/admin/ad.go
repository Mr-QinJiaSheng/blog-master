package admin

import (
	"errors"
	models "go-blog/models/admin"
	"go-blog/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

// AdController operations for Ad
type AdController struct {
	BaseController
}

// URLMapping ...
func (c *AdController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Ad
// @Param	body		body 	models.Ad	true		"body for Ad content"
// @Success 201 {int} models.Ad
// @Failure 403 body is empty
// @router /ad [post]
func (c *AdController) Post() {
	response := make(map[string]interface{})
	gid := c.GetString("gid")
	group := c.GetString("group")
	title := c.GetString("title")
	image := c.GetString("image")
	url := c.GetString("url")

	if gid == "" {

		worker, err := utils.NewIdWorker(1)
		if err != nil {
			response["msg"] = err.Error()
			response["code"] = 500
			response["err"] = err.Error()
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}

		uid, err := worker.GetNextId()
		if err != nil {
			response["msg"] = err.Error()
			response["code"] = 500
			response["err"] = err.Error()
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
		uidstr := strconv.FormatInt(uid, 10)
		gid = "GID" + uidstr
	}

	if len(group) == 0 || len(title) == 0 || len(image) == 0 || len(url) == 0 {
		response["msg"] = "参数有误，必填项不能为空！"
		response["code"] = 500
		response["err"] = "参数有误，必填项不能为空！"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	if _, err := models.AddAd(&models.Ad{
		Title: title,
		Image: image,
		Url:   url,
		Gid:   gid,
		Group: group,
	}); err == nil {
		response["msg"] = "OK"
		response["code"] = 200
		c.Ctx.Output.SetStatus(201)
	} else {
		response["msg"] = err.Error()
		response["code"] = 500
		response["err"] = err.Error()
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *AdController) Add() {

	type GData struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	}

	o := orm.NewOrm()

	group := []*models.Ad{}
	o.QueryTable(new(models.Ad)).GroupBy("gid").All(&group)
	//fmt.Printf("%v", group)
	var g []GData
	for _, v := range group {
		g = append(g, GData{
			Value: v.Gid,
			Name:  v.Group,
		})
	}

	if len(g) == 0 {
		c.Data["json"] = "[]"
	} else {
		c.Data["json"] = g
	}

	// c.Data["json"] = string(data)
	// c.ServeJSON()
	// c.StopRun()
	c.TplName = "admin/ad-add.html"
}

func (c *AdController) Edit() {

	o := orm.NewOrm()

	type GData struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	}

	group := []*models.Ad{}
	o.QueryTable(new(models.Ad)).GroupBy("gid").All(&group)
	//fmt.Printf("%v", group)
	var g []GData
	for _, v := range group {
		g = append(g, GData{
			Value: v.Gid,
			Name:  v.Group,
		})
	}

	c.Data["json"] = g

	id, _ := c.GetInt("id")
	v, err := models.GetAdById(id)
	if err != nil {
		c.Abort("非法操作！")
	}
	c.Data["Data"] = v
	c.TplName = "admin/ad-edit.html"
}

// GetOne ...
// @Title Get One
// @Description get Ad by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ad
// @Failure 403 :id is empty
// @router /ad/:id [get]
func (c *AdController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAdById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Ad
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Ad
// @Failure 403
// @router /ad [get]
func (c *AdController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	//var limit int64 = 10
	//var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				//c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.Abort(errors.New("Error: invalid query key/value pair").Error())
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAd(query, fields, sortby, order, 0, 0)
	if err != nil {
		c.Abort(err.Error())
	}
	type Data struct {
		Gid   string
		Group string
		Ad    []*models.Ad
	}

	var data []*Data

	for _, value := range l {
		ad := value.(models.Ad)
		var flag bool = false
		for _, v := range data {
			if v.Gid == ad.Gid {
				v.Ad = append(v.Ad, &ad)
				flag = true
				break
			}
		}

		if flag == true {
			continue
		}

		data = append(data, &Data{
			Gid:   ad.Gid,
			Group: ad.Group,
			Ad:    []*models.Ad{&ad},
		})
	}

	c.Data["Data"] = data
	//c.ServeJSON()
	c.TplName = "admin/ad.html"
}

// Put ...
// @Title Put
// @Description update the Ad
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Ad	true		"body for Ad content"
// @Success 200 {object} models.Ad
// @Failure 403 :id is not int
// @router /ad/:id [put]
func (c *AdController) Put() {
	response := make(map[string]interface{})
	gid := c.GetString("gid")
	group := c.GetString("group")
	title := c.GetString("title")
	image := c.GetString("image")
	url := c.GetString("url")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if gid == "" {

		worker, err := utils.NewIdWorker(1)
		if err != nil {
			response["msg"] = err.Error()
			response["code"] = 500
			response["err"] = err.Error()
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}

		uid, err := worker.GetNextId()
		if err != nil {
			response["msg"] = err.Error()
			response["code"] = 500
			response["err"] = err.Error()
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
		uidstr := strconv.FormatInt(uid, 10)
		gid = "GID" + uidstr
	}

	if len(group) == 0 || len(title) == 0 || len(image) == 0 || len(url) == 0 {
		response["msg"] = "参数有误，必填项不能为空！"
		response["code"] = 500
		response["err"] = "参数有误，必填项不能为空！"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	if err := models.UpdateAdById(&models.Ad{
		Id:    id,
		Title: title,
		Image: image,
		Url:   url,
		Gid:   gid,
		Group: group,
	}); err == nil {
		response["msg"] = "OK"
		response["code"] = 200
		c.Ctx.Output.SetStatus(201)
	} else {
		response["msg"] = err.Error()
		response["code"] = 500
		response["err"] = err.Error()
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Ad
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /ad/:id [delete]
func (c *AdController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAd(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *AdController) DeleteGroup() {
	gid := c.GetString("gid")
	response := make(map[string]interface{})
	o := orm.NewOrm()
	_, err := o.Delete(&models.Ad{Gid: gid}, "gid")

	if err != nil {
		response["msg"] = err.Error()
		response["code"] = 500
		response["err"] = err.Error()

	} else {
		response["msg"] = "OK"
		response["code"] = 200
		c.Ctx.Output.SetStatus(201)
	}
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
