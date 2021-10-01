package routers

import (
	"go-blog/controllers/admin"
	"go-blog/controllers/common"
	"go-blog/controllers/home"
	"go-blog/controllers/installer"
	"go-blog/controllers/wechat"

	"github.com/astaxie/beego"
)

func init() {

	// 安装功能
	// beego.InsertFilter("*",beego.BeforeExec,filter.Installer)

	adminNs := beego.NewNamespace("/admin",
		beego.NSInclude(
			&admin.MenuController{},
			&admin.LinkController{},
			&admin.FileController{},
			&admin.CronController{},
			&admin.CustomerController{},
		),

		beego.NSRouter("/get-review", &admin.CronController{}, "get:GetReview"),

		beego.NSRouter("/user", &admin.UserController{}, "get:List;post:Save"),
		beego.NSRouter("/user/edit", &admin.UserController{}, "get:Put"),
		beego.NSRouter("/user/status", &admin.UserController{}, "Post:Delete"),
		beego.NSRouter("/user/update", &admin.UserController{}, "Post:Update"),
		beego.NSRouter("/user/add", &admin.UserController{}, "get:Add"),

		// 站点设置
		beego.NSRouter("/setting", &admin.SettingController{}, "get:Add"),
		beego.NSRouter("/setting/save", &admin.SettingController{}, "post:Save"),
		beego.NSRouter("/notice", &admin.SettingController{}, "get:Notice"),
		beego.NSRouter("/notice/save", &admin.SettingController{}, "post:NoticeSave"),
		beego.NSRouter("/about", &admin.SettingController{}, "get:About"),
		beego.NSRouter("/about/save", &admin.SettingController{}, "post:AboutSave"),

		// 后台文章模块
		beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),
		beego.NSRouter("/article", &admin.ArticleController{}, "get:List;post:Save"),
		beego.NSRouter("/article/edit", &admin.ArticleController{}, "get:Put"),
		beego.NSRouter("/article/delete", &admin.ArticleController{}, "Post:Delete"),
		beego.NSRouter("/article/update", &admin.ArticleController{}, "Post:Update"),
		beego.NSRouter("/article/add", &admin.ArticleController{}, "get:Add"),
		beego.NSRouter("/article/top", &admin.ArticleController{}, "Post:Top"),
		beego.NSRouter("/article/get", &admin.ArticleResourcesController{}, "Post:GetArticle"),

		beego.NSRouter("/article/cron/get", &admin.ArticleResourcesController{}, "Get:GetCron"),

		// 后台分类模块
		beego.NSRouter("/cate", &admin.CateController{}, "get:List;post:Save"),
		beego.NSRouter("/cate/add", &admin.CateController{}, "get:Add"),
		beego.NSRouter("/cate/edit", &admin.CateController{}, "get:Put"),
		beego.NSRouter("/cate/delete", &admin.CateController{}, "Post:Delete"),
		beego.NSRouter("/cate/update", &admin.CateController{}, "Post:Update"),
		// 后台登录
		beego.NSRouter("/login", &admin.LoginController{}, "Get:Sign;Post:Login"),

		// 后台评论
		beego.NSRouter("/review", &admin.ReviewController{}, "get:List"),
		beego.NSRouter("/review/edit", &admin.ReviewController{}, "Get:Put"),
		beego.NSRouter("/review/delete", &admin.ReviewController{}, "Post:Delete"),
		beego.NSRouter("/review/update", &admin.ReviewController{}, "Post:Update"),

		// 后台留言
		beego.NSRouter("/message", &admin.MessageController{}, "get:List"),

		// 后台留言模块
		beego.NSRouter("/message/update", &admin.MessageController{}, "Post:Update"),
		beego.NSRouter("/message/edit", &admin.MessageController{}, "Get:Put"),
		beego.NSRouter("/message/delete", &admin.MessageController{}, "Post:Delete"),
	)

	beego.AddNamespace(adminNs)
	// 公众号
	beego.Router("/wechat", &wechat.MainController{}, "Get:CheckToken;Post:Hello")
	beego.Router("/wechat/create/menu", &wechat.MenuController{}, "Get:CreateMenu;Post:CreateMenu")
	beego.Router("/wechat/user/get", &wechat.UserController{}, "Get:GetUser")
	beego.Router("/wechat/user/list", &wechat.UserController{}, "Get:List")
	beego.Router("/wechat/addnews", &wechat.MaterialController{}, "Get:AddNews")

	beego.Router("/", &home.MainController{})
	beego.Router("/admin", &admin.MainController{}, "get:Index")
	// 前台列表
	beego.Router("/list.html", &home.ArticleController{}, "get:List")
	// 前台详情
	beego.Router("/detail/:id([0-9]+).html", &home.ArticleController{}, "get:Detail")
	// 前台统计文章PV
	beego.Router("/pv/:id([0-9]+).html", &home.ArticleController{}, "get:Pv")
	// 前台留言列表
	beego.Router("/message.html", &home.MessageController{}, "get:Get")
	// 前台保存
	beego.Router("/message/save", &home.MessageController{}, "Post:Save")

	// 评论保存
	beego.Router("article/review", &home.ArticleController{}, "Post:Review")
	beego.Router("article/review/:id([0-9]+).html", &home.ArticleController{}, "Get:ReviewList")
	beego.Router("article/like", &home.ArticleController{}, "Post:Like")

	// 文件上传
	beego.Router("/uploads.html", &common.UploadsController{}, "Post:Uploads")

	// 安装
	beego.Router("/installer", &installer.InstallController{}, "Get:CheckEnv")
	beego.Router("/installer/create", &installer.InstallController{}, "Get:Install")

}
