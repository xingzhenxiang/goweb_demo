package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"web_demo/framework"
	"web_demo/service"
	"web_demo/utils"
)

/**
 * r.PostFormValue  : 可以解析 Post/PUT Content-Type=application/x-www-form-urlencoded 或 Content-Type=multipart/form-data
 */

type TopicConterller struct {
	//Data map[interface{}]interface{}
}

var TopicService = new(service.TopicService)

//POST Content-Type=application/x-www-form-urlencoded
func (p *TopicConterller) addtopic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		cookie, err := r.Cookie("testcookiename")
		if err != nil {
			fmt.Println(err)

			http.Redirect(w, r, "/login", 302)
			return
		}
		fmt.Println(cookie)
		t, _ := template.ParseFiles("view/addtopic.html")
		log.Println(t.Execute(w, nil))
	} else {
		title := r.PostFormValue("title")
		topicspreview := r.PostFormValue("topicspreview")
		topicsinfo := r.PostFormValue("topicsinfo")

		title = strings.ToLower(title)
		if utils.Empty(title) || utils.Empty(topicspreview) || utils.Empty(topicsinfo) {
			framework.ResultFail(w, "title, topicspreview, topicsinfo can not be empty")
			return
		}
		id := TopicService.Insert(title, topicspreview, topicsinfo)
		if id <= 0 {
			framework.ResultFail(w, "addtopic fail")
			return
		}
		framework.ResultOk(w, "addtopic success")
	}
}

//POST Content-Type=application/x-www-form-urlencoded

// GET/POST
func (p *TopicConterller) topicFindAll(w http.ResponseWriter, r *http.Request) {
	Topics := TopicService.SelectAllTopic()
	htmlname := "view/topics.html"
	framework.ResultTmpOk(w, Topics, htmlname)
}

func (p *TopicConterller) topicDeleteById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("testcookiename")
	if err != nil {
		fmt.Println(err)

		http.Redirect(w, r, "/login", 302)
		return
	}
	fmt.Println(cookie)
	urlstr := r.URL.String()
	n := strings.Index(urlstr, "=")
	TopicService.TopicDeleteById(string(urlstr[n+1:]))
	http.Redirect(w, r, "/index", 302)
}

func (p *TopicConterller) detailTopic(w http.ResponseWriter, r *http.Request) {
	urlstr := r.URL.String()
	n := strings.Index(urlstr, "=")
	Topics := TopicService.SelectTopicById(string(urlstr[n+1:]))
	//fmt.Println(string(urlstr[n+1:]))
	htmlname := "view/detailtopic.html"

	framework.ResultTmpOk(w, Topics, htmlname)
}

func (p *TopicConterller) modifyTopic(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		cookie, err := r.Cookie("testcookiename")
		if err != nil {
			fmt.Println(err)

			http.Redirect(w, r, "/login", 302)
			return
		}
		fmt.Println(cookie)
		urlstr := r.URL.String()
		n := strings.Index(urlstr, "=")
		Topics := TopicService.SelectTopicById(string(urlstr[n+1:]))
		fmt.Println(string(urlstr[n+1:]))
		htmlname := "view/modifytopic.html"

		framework.ResultTmpOk(w, Topics, htmlname)

		//t, _ := template.ParseFiles("view/modifytopic.html")
		//log.Println(t.Execute(w, nil))
	} else {
		id := r.PostFormValue("id")

		title := r.PostFormValue("title")
		topicspreview := r.PostFormValue("topicspreview")
		topicsinfo := r.PostFormValue("topicsinfo")

		if utils.Empty(title) || utils.Empty(topicspreview) || utils.Empty(topicsinfo) {
			framework.ResultFail(w, "title, topicspreview, topicsinfo can not be empty")
			return
		}
		num := TopicService.TopicModifyB(title, topicspreview, topicsinfo, id)
		if num <= 0 {
			framework.ResultFail(w, "addtopic fail")
			return
		}
		framework.ResultOk(w, "addtopic success")
	}
}

func (p *TopicConterller) listTopic(w http.ResponseWriter, r *http.Request) {
	urlstr := r.URL.String()
	n := strings.Index(urlstr, "=")
	strnum := string(urlstr[n+1:])
	page, error := strconv.Atoi(strnum)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
		page = 1
	}
	pre_page := 3
	Topics := TopicService.LimitList(pre_page, page)
	totals := TopicService.GetDataNum()
	res := utils.Paginator(page, pre_page, totals)
	var Data map[interface{}]interface{}
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	Data = make(map[interface{}]interface{})
	Data["datas"] = Topics  //博文的数据
	Data["paginator"] = res //分页的数据
	Data["totals"] = totals //分页的数据
	htmlname := "view/list.html"

	framework.ResultTmpOk(w, Data, htmlname)
}
