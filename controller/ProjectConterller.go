package controller

import (
	"web_demo/framework"

	"github.com/gorilla/sessions"
)

type ProjectConterller struct {
	ic IndexConterller
	uc UserConterller
	tc TopicConterller
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

const (
	authenticated string = "authenticated"
	cookieName    string = "cookie-name"
)

func (p *ProjectConterller) Router(router *framework.RouterHandler) {
	router.Router("/", p.tc.listTopic)                   //js 伪分页//替换为真正分页
	router.Router("/index", p.tc.listTopic)              ///js 伪分页/替换为真正分页
	router.Router("/topic", p.tc.topicFindAll)           //js 伪分页
	router.Router("/register", p.uc.register)            //用户注册
	router.Router("/login", p.uc.login)                  //用户登录
	router.Router("/findAll", p.uc.findAll)              //所有用户
	router.Router("/topic/delete", p.tc.topicDeleteById) // 博文删除
	router.Router("/topic/add", p.tc.addtopic)           // 博文添加
	router.Router("/topic/modify", p.tc.modifyTopic)     // 博文修改
	router.Router("/topic/detail", p.tc.detailTopic)     // 博文详细内容
	router.Router("/List", p.tc.listTopic)               //真正分页

}
