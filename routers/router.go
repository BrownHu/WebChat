package routers

import (
	"LetsChat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.ChatController{},"get:Join")  //default page
    beego.Router("/", &controllers.ChatController{},"post:Verify")  //添加新聊者 ajax 返回json
    beego.Router("/chat", &controllers.ChatController{},"get:Index")  //进入聊天
    beego.Router("/chat", &controllers.ChatController{},"post:Chat")  //发消息
    beego.Router("/refresh", &controllers.ChatController{},"get:Refresh")  //发消息
    beego.Router("/offline", &controllers.ChatController{},"post:Offline")  //离线

}
