package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"web_demo/framework"
	"web_demo/service"
	"web_demo/utils"
)

/**
 * r.PostFormValue  : 可以解析 Post/PUT Content-Type=application/x-www-form-urlencoded 或 Content-Type=multipart/form-data
 */

type UserConterller struct {
}

var userService = new(service.UserService)

/*
func (p *UserConterller) Router(router *framework.RouterHandler) {
	router.Router("/register", p.register)
	router.Router("/login", p.login)
	router.Router("/findAll", p.findAll)

}*/

//POST Content-Type=application/x-www-form-urlencoded
func (p *UserConterller) register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view/register.html")
		log.Println(t.Execute(w, nil))
	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		username = strings.ToLower(username)
		if utils.Empty(username) || utils.Empty(password) {
			framework.ResultFail(w, "username or password can not be empty")
			return
		}
		id := userService.Insert(username, password)
		if id <= 0 {
			framework.ResultFail(w, "register fail")
			return
		}
		framework.ResultOk(w, "register success")
	}
}

// GET/POST
func (p *UserConterller) findAll(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("testcookiename")
	if err != nil {
		fmt.Println(err)

		http.Redirect(w, r, "/login", 302)
		return
	}
	fmt.Println(cookie)

	users := userService.SelectAllUser()
	htmlname := "view/findall.html"
	framework.ResultTmpOk(w, users, htmlname)
}

func (p *UserConterller) login(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		if r.URL.String() == "/login?exit=true" {
			cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
			http.SetCookie(w, &cookie)
		}

		t, _ := template.ParseFiles("view/login.html")
		log.Println(t.Execute(w, nil))
	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		username = strings.ToLower(username)

		if utils.Empty(username) || utils.Empty(password) {
			framework.ResultFail(w, "username or password can not be empty")
			return
		}
		users := userService.SelectUserByName(username)
		if len(users) == 0 {
			framework.ResultFail(w, "user does not exist")
			return
		}
		if users[0].Password != password {
			framework.ResultFail(w, "password error")
			return
		}

		cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", MaxAge: 86400}
		http.SetCookie(w, &cookie)
		framework.ResultOk(w, "login success") //这个操作在cookie操作前，会导致cookie写入失败

	}
}
