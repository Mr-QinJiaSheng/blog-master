package home

import (
	"fmt"
	"go-blog/models/admin"
	"go-blog/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

func (ctl *LoginController) Sign() {
	ctl.TplName = "home/" + ctl.Template + "/login.html"
}
func (ctl *LoginController) Login() {

	response := make(map[string]interface{})

	username := ctl.GetString("username")
	password := ctl.GetString("password")

	//fmt.Println(username)
	//fmt.Println(password)
	if username == "" || password == "" {
		response["code"] = 500
		response["msg"] = "用户名或者密码不能为空！"
		ctl.Data["json"] = response
		ctl.ServeJSON()
		ctl.StopRun()
	}

	password = utils.PasswordMD5(password, "nihongdengxia.com")

	if customer, ok := admin.CustomerLogin(username, password); ok {
		ctl.SetSession("Customer", *customer)
		response["code"] = 200
		response["msg"] = "登录成功！"
		response["referer"] = ctl.Ctx.Request.Referer()
		//ctl.Redirect(ctl.Ctx.Request.Referer(), 302)
	} else {
		response["code"] = 500
		response["msg"] = "登录失败！"
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()

}

func (ctl *LoginController) GetCaptcha() {
	response := make(map[string]interface{})
	var captcha string
	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)

	email := ctl.GetString("email")

	if email == "" {
		response["code"] = 500
		response["msg"] = "邮箱不能为空！"
		ctl.Data["json"] = response
		ctl.ServeJSON()
		ctl.StopRun()
	}

	if err != nil {
		response["code"] = 500
		response["msg"] = "服务器出错了！"
		ctl.Data["json"] = response
		ctl.ServeJSON()
		ctl.StopRun()
	}

	o := orm.NewOrm()
	exist := o.QueryTable("customer").Filter("Email", email).Exist()
	if exist {
		response["msg"] = "该邮箱已经存在！"
		response["code"] = 500
		response["err"] = ""
		ctl.Data["json"] = response
		ctl.ServeJSON()
		ctl.StopRun()
	}

	//bm.Put("astaxie", 1, 10*time.Second)
	//captcha = bm.Get("reg-captcha").(string)
	//if captcha == "" {
	captcha = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	bm.Put("reg-captcha", captcha, 5*time.Minute)
	//}
	//bm.IsExist("astaxie")
	//bm.Delete("astaxie")

	//captcha = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	flag := utils.SendEmail(utils.Email{
		From: "1920853199@qq.com",
		To:   email,
		Header: map[string]string{
			"Subject": "霓虹灯下注册验证码",
		},
		Template: "views/home/nihongdengxia/reg-email.html",
		Data: utils.Reg{
			//Tag:  "注册",
			Code: captcha,
			//Date: time.Now().Format("2006-01-02 15:04:05"),
		},
	})
	if flag == nil {
		response["code"] = 200
		response["msg"] = "发送成功！"
	} else {
		response["code"] = 500
		response["msg"] = flag.Error()
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()

}

func (c *LoginController) Regist() {
	email := c.GetString("email")
	pass := c.GetString("pass")
	repass := c.GetString("repass")
	code := c.GetString("code")

	response := make(map[string]interface{})

	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)

	if err != nil || !bm.IsExist("reg-captcha") {
		response["msg"] = "服务器出错！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	captcha := bm.Get("reg-captcha").(string)
	if captcha != code {
		response["msg"] = "验证码错误！"
		response["code"] = 500
		response["err"] = "验证码错误！"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	valid := validation.Validation{}
	valid.Email(email, "email")
	valid.Required(pass, "pass")
	valid.Required(pass, "repass")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message)
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = err.Key + " " + err.Message
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
	}

	if pass != repass {
		response["msg"] = "密码不一致！"
		response["code"] = 500
		response["err"] = "密码不一致！"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	o := orm.NewOrm()
	exist := o.QueryTable("customer").Filter("Email", email).Exist()
	if exist {
		response["msg"] = "该邮箱已经存在！"
		response["code"] = 500
		response["err"] = ""
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	//var id int64
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
	if id, err := admin.AddCustomer(&admin.Customer{
		Uid:      uidstr,
		Username: utils.CreateUsername(),
		Email:    email,
		Password: utils.PasswordMD5(pass, "nihongdengxia.com"),
		Status:   1,
		Integral: 20,
		Fans:     0,
		Focus:    0,
		Created:  time.Now(),
		Updated:  time.Now(),
	}); err == nil {

		o := orm.NewOrm()
		customer := admin.Customer{Id: int(id)}
		o.Read(&customer)
		c.SetSession("Customer", customer)

		bm.Delete("reg-captcha")
		response["msg"] = "注册成功！"
		response["code"] = 200
		response["id"] = id
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
