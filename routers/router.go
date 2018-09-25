package routers

import (
	"BeeGo-RabbitConnect/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")
    beego.Router("/sendMess", &controllers.MainController{}, "get:MessageCatcher")
}
