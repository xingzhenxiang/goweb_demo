package framework

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func ResultOk(w http.ResponseWriter, data string) {
	io.WriteString(w, data)
}

func ResultFail(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

func ResultJsonOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(data)
	w.Write(json)
}

func ResultTmpOk(w http.ResponseWriter, data interface{}, htmlname string) {
	//w.Header().Set("Content-Type", "application/json")
	//fmt.Println(htmlname)
	t, err := template.ParseFiles(htmlname)
	if err != nil {
		fmt.Println(err)
	}
	//json, _ := json.Marshal(data)
	//w.Write(json)
	//fmt.Println(data)
	//err = t.ExecuteTemplate(w, htmlname, data)
	t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func OpResult(w http.ResponseWriter, data string) {
	var Data map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	Data = make(map[string]string)
	Data["Message"] = data + "  3秒后自动跳转到首页"
	htmlname := "view/sucess.html"

	ResultTmpOk(w, Data, htmlname)
}
