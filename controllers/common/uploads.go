package common

import (
	"go-blog/utils"
	"math/rand"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type UploadsController struct {
	beego.Controller
}

func (c *UploadsController) Uploads() {

	response := make(map[string]interface{})

	f, h, err := c.GetFile("editormd-image-file")
	defer f.Close()
	if err != nil {
		response["message"] = err.Error()
		response["success"] = 0
	} else {

		ext := path.Ext(h.Filename)
		filename := time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(1000)) + ext

		filepath := "static/uploads/" + time.Now().Format("200601")
		if err = utils.CheckDir(filepath); err != nil {
			response["message"] = err.Error()
			response["success"] = 0
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
		path := filepath + "/" + filename
		err := c.SaveToFile("editormd-image-file", path)
		if err != nil {
			response["message"] = err.Error()
			response["success"] = 0
		} else {
			response["success"] = 1
			response["message"] = "Success."
			response["url"] = "/" + path
		}
	}

	c.Data["json"] = response
	c.ServeJSON()

}
