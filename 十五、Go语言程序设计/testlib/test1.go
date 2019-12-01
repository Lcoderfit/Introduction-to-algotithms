package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/astaxie/beego"
	"net/rpc"
	"encoding/json"
)

func main() {
	str := "&I&love&coding"
	res := strings.Split(str, "&")
	fmt.Println("result:", res)
	fmt.Println("len:", len(res))
	fmt.Println("cap:", cap(res))
	fmt.Printf("%T", res)
	for _, s := range res {
		fmt.Println(s)
	}
	log.Println("11/30")
}

func (this *baseRouter) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	lang := this.Input().Get("lang")

	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	} else {
		isNeedRedir = true
	}

	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5]
			if !i18n.Exist(al) {
				lang = al
			}
		}
	}

	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

}

type AddController struct {
	beego.Controller
}

func (this *AddController) Get() {
	this.Data["content"] = "value"
	this.Layout = "admin/layout.html"
	this.TplName = "admin/add.tpl"
}

func (this *AddController) Post() {
	pkgname := this.GetString("pkgname")
	content := this.GetString("content")
	pk := models.GetCruPkg(pkgname)
	if pk.Id == 0 {
		var pp models.PkgEntity
		pp.Pid = 0
		pp.Pathname = pkgname
		models.InsertPkg(pp)
		pk = models.GetCruPkg(pkgname)
	}

	var at models.Article
	at.Pkgid = pk.id
	at.Content = content
	models.InsertArticle(at)
	this.Ctx.Redirect(302, "/admin/index")
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["json"] = map[string]interface{}{"name": "astaxie"}
	this.ServeJSON()
	this.StopRun() // 终止运行，不再执行后面的Get、POST等其他请求方法
}

type BaseAdminRouter struct {
	baseController
}

func (this *BaseAdminRouter) NestPrepare() {
	if this.CheckActiveRedirect() {
		return
	}

	if !this.user.IsAdmin {
		models.LogoutUser(&this.Controller)

		this.FlashWrite("NotPermit", "true")

		this.Redirect("/login", 302)
		return
	}

	this.Data["IsAdmin"] = true

	if app, ok := this.AppController.(ModelPreparer); ok {
		app.ModelPrepare()
		return
	}
}

func (this *BaseAdminRouter) Get() {
	this.TplName = "Get.tpl"
}

func d() {
	beego.Get("/", func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})
	
	beego.Post("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	
	beego.Any("/foo", func(ctx *content.Context) {
		ctx.Output.Body([]byte("bar"))
	})
	
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("FUCKER"))
	}))
	
	beego.Post("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("asdf"))
	})
	
	beego.Put("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("lcoder"))
	})
	
	beego.Patch("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("share"))
	})
	
	beego.Head("/", func(ctx *context.Context){
		ctx.Output.Body([]byte("fit"))
	})
	
	beego.Options("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("fuc"))
	})
	
	beego.Delete("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("delete"))
	})
	
	beego.Any("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Any"))
	})
}

func d() {
	// 自定义handler实现
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	beego.Handler("/rpc", s)
}

func e() {
	// 固定路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.UserController{})
	beego.Router("/admin/index", &admin.ArticleController{})
	beego.Router("/admin/addpkg", &admin.AddController{})
}

func f() {
	// 正则路由
	// /api/123	/api/
	beego.Router("api/?:id", &controllers.RController{})
	// /api/12	但不能匹配/api/
	beego.Router("/api/:id", &controllers.RController{})
	// /api/123 但不能匹配/api/lco
	beego.Router("/api/:id([0-9]+)", &controllers.RController{})
	// /user/astaxie 此时:username为astaxie
	beego.Router("/user/:username([\\w]+)")
	// /download/file/api.xml
	beego.Router("/download/*.*", &controllers.RController{})
	
	// :id为int类型，效果与([0-9]+)相同
	beego.Router("/:id:int", &controllers.RController{})
	// :hi为string类型，效果与([\\w]+)类似
	beego.Router("/:hi:string", &controllers.RController{})
	// :id为正则类型，匹配/cms_123.html
	beego.Router("/cms_:id([0-9]).html", &controllers.CmsController{})
}

func g() {
	// 获取上面的变量
	this.Ctx.Input.Param(":id")
	this.Ctx.Input.Param(":username")
}

func h() {
	// 自定义方法及restful规则, httpmethod:funcname
	beego.Router("/api/list", &RestController{}, "*:ListFood")
	beego.Router("/api/create", &RestController{}, "post:CreateFood")
	beego.Router("/api/update", &RestController{}, "put:UpdateFood")
	beego.Router("/api/delete", &RestController{}, "delete:DeleteFood")
	
	// 多个HTTPMethod对应一个函数
	beego.Router("/api", &RestController{}, "get, post:ApiFunc")
	// 不同的HTTPMethod对应多个函数
	beego.Router("/simple", &RestController{}, "get:GetFunc;post:PostFunc")
	
	// 如果同时有*:AllFunc;post:PostFunc; 那么执行post:PostFunc,而不执行AllFunc
	beego.Router("/sample", &SimpleController{}, "*:AllFunc;post:PostFunc")
}

func i() {
	// 自动匹配beego
	// /object/login:匹配ObjectController的login方法
	// /object/logout:匹配ObjectController的logout方法
	// /object/blog/2019/1/1:自动调用ObjectController的Blog方法，剩下的参数：map[0:2019 1:1 2:1]
	// /object/LCODER: beego会自动转化为小写，调用ObjectController的Lcoder方法
	beego.AutoRouter(&controllers.ObjectController{})
}

// 注解路由
type CMSController struct {
	beego.Controller 
}

func (c *CMSController) URLMapping() {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
// 类似于：beego.Router("/staticblock/:key", &CMSController{}, "get:StaticBlock")
func (this *CMSController) StaticBlock() {
	
}

// @router /all/:key [get]
// 编写注解路由之后，在router.go中添加：beego.Include(&CMSController{})
func (this *CMSController) AllBlock() {

}


// 命名空间
//初始化 namespace
func k() {
	ns :=
	beego.NewNamespace("/v1",
		// 支持满足的条件就执行namespace，不满足就不执行
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "api.beego.me" {
				return true
			}
			return false
		}),
		beego.NSBefore(auth),
		beego.NSGet("/notallowed", func(ctx *context.Context) {
			ctx.Output.Body([]byte("notAllowed"))
		}),
		beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
		beego.NSRouter("changepassword", &UserController{}),
		beego.NSNamespace("/shop",
			beego.NSBefore(sentry),
			beego.NSGet("/:id", func(ctx *context.Context) {
				ctx.Output.Body([]byte("notAllowed"))
			}),
		),
		beego.NSNamespace("/cms",
			beego.NSInclude(
				&controllers.MainController{},
				&controllers.CMSController{},
				&controllers.BlockController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

func l() {
	// 嵌套namespace
	ns :=
	beego.NewNamespace("/v1",
		beego.NSGet("/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte("shopinfo"))
		}),
	),
	beego.NSNamespace("/order",
		beego.NSGet("/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte("orderinfo"))
		})
	),
	beego.NSNamespace("/crm",
		beego.NSGet("/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte("crminfo"))
		}),
	)
}

func m() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSCond(func(ctx *context.Context) bool {
				if ua := ctx.Input.Request.UserAgent(); ua != "" {
					return true
				}
				return false
			}),
			beego.NSNamespace()
		)
}

// ns :=
//     beego.NewNamespace("/api",
//         //此处正式版时改为验证加密请求
//         beego.NSCond(func(ctx *context.Context) bool {
//             if ua := ctx.Input.Request.UserAgent(); ua != "" {
//                 return true
//             }
//             return false
//         }),
//         beego.NSNamespace("/ios",
//             //CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
//             beego.NSNamespace("/create",
//                 // /api/ios/create/node/
//                 beego.NSRouter("/node", &apis.CreateNodeHandler{}),
//                 // /api/ios/create/topic/
//                 beego.NSRouter("/topic", &apis.CreateTopicHandler{}),
//             ),
//             beego.NSNamespace("/read",
//                 beego.NSRouter("/node", &apis.ReadNodeHandler{}),
//                 beego.NSRouter("/topic", &apis.ReadTopicHandler{}),
//             ),
//             beego.NSNamespace("/update",
//                 beego.NSRouter("/node", &apis.UpdateNodeHandler{}),
//                 beego.NSRouter("/topic", &apis.UpdateTopicHandler{}),
//             ),
//             beego.NSNamespace("/delete",
//                 beego.NSRouter("/node", &apis.DeleteNodeHandler{}),
//                 beego.NSRouter("/topic", &apis.DeleteTopicHandler{}),
//             )
//         ),
//     )

// beego.AddNamespace(ns)