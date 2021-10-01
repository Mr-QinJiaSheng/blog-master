package admin

import (
	"errors"
	"go-blog/models/admin"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// CronController operations for Cron
type CronController struct {
	beego.Controller
}

// URLMapping ...
func (c *CronController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Cron
// @Param	body		body 	admin.Cron	true		"body for Cron content"
// @Success 201 {int} admin.Cron
// @Failure 403 body is empty
// @router cron/save [post]
func (c *CronController) Post() {

	response := make(map[string]interface{})

	if _, err := admin.AddCron(&admin.Cron{
		Title:  c.GetString("title"),
		Cron:   c.GetString("cron"),
		Url:    c.GetString("url"),
		Status: 1,
	}); err == nil {
		response["msg"] = "OK"
		response["code"] = 200
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = response
	} else {
		response["msg"] = "Fields"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Cron by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} admin.Cron
// @Failure 403 :id is empty
// @router cron/:id [get]
func (c *CronController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := admin.GetCronById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Cron
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} admin.Cron
// @Failure 403
// @router cron [get]
func (c *CronController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
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
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, _ := admin.GetAllCron(query, fields, sortby, order, offset, limit)
	c.Data["Data"] = l
	c.TplName = "admin/article-get.html"

}

// Put ...
// @Title Put
// @Description update the Cron
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	admin.Cron	true		"body for Cron content"
// @Success 200 {object} admin.Cron
// @Failure 403 :id is not int
// @router cron/:id [put]
func (c *CronController) Put() {
	response := make(map[string]interface{})

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := admin.Cron{Id: id}

	cron, err := admin.GetCronById(id)
	if err != nil {
		response["msg"] = "field"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
		c.ServeJSON()
	}
	if cron.Status == 1 {
		v.Status = 2
	} else {
		v.Status = 1
	}

	v.Title = cron.Title
	v.Cron = cron.Cron
	v.Url = cron.Url
	v.Created = cron.Created

	if err := admin.UpdateCronById(&v); err == nil {
		response["msg"] = "OK"
		response["code"] = 200
		c.Data["json"] = response
	} else {
		response["msg"] = "field"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Cron
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router cron/:id [delete]
func (c *CronController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := admin.DeleteCron(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *CronController) GetReview() {
	c.TplName = "admin/review-get.html"
}
