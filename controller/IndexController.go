package controller

import (
	"net/http"
	"web_demo/framework"
)

type IndexConterller struct {
}

/*
func (p *IndexConterller) Router(router *framework.RouterHandler) {
	router.Router("/index", p.index)
	router.Router("/", p.index)

}*/
func (p *IndexConterller) index(w http.ResponseWriter, r *http.Request) {
	//	users := userService.SelectAllUser()
	framework.ResultJsonOk(w, "kehl's index")
}
