package common

import (
	"errors"
	models "go-blog/models/admin"
	"go-blog/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MusicController struct {
	beego.Controller
}

func (c *MusicController) Get() {

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

	l, err := models.GetAllMusic(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()

}

func (c *MusicController) Update() {

	limit, _ := c.GetInt("limit", 10) // 页数
	token := c.GetString("token")
	start, _ := c.GetInt("start", 0)

	if token != "zmkmmmmmh" {
		c.Data["json"] = "非法操作！"
		c.ServeJSON()
		c.StopRun()
	}
	var m models.Music
	if start == 0 {
		o := orm.NewOrm()
		err := o.QueryTable(new(models.Music)).OrderBy("-id").One(&m)

		if err != nil {
			if err == orm.ErrNoRows {
				m.SongId = 100000
			} else {
				c.Data["json"] = err.Error()
				c.ServeJSON()
				c.StopRun()
			}
		}
	} else {
		m.SongId = start
	}
	//*/1 0 * * *
	// 0 0 * * *
	id := m.SongId
	//id := 100000
	for i := 0; i < limit; i++ {
		id = utils.MusicSave(id)
		//time.Sleep(time.Second)
	}

	//fmt.Println("更新10条曲库！")
	c.Data["json"] = "更新" + strconv.Itoa(limit) + "条曲库！"
	c.ServeJSON()
	c.StopRun()
}
