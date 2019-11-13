package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/unrolled/render"
)

//登陆界面
func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("method:", req.Method) //获取请求的方法
		if req.Method == "GET" {           //GET 方法请求初始登陆界面
			formatter.HTML(w, http.StatusOK, "login", struct{}{})
		} else if req.Method == "POST" { //POST 方法请求登陆数据，执行登陆操作
			req.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
			//注意:如果没有调用ParseForm方法，下面无法获取表单的数据

			log.Println("username:", req.Form["username"])

			//获取主页
			//使用 formatter 的 HTML 直接将数据注入模板，并输出到浏览器
			formatter.HTML(w, http.StatusOK, "index", struct {
				ID string `json:"id"`
			}{ID: strings.Join(req.Form["username"], " ")})
		}
	}
}
