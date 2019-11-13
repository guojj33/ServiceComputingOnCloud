package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "templates",       //模板所在的文件夹
		Extensions: []string{".html"}, //模板文件后缀
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter() //路由器

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	//获取根目录路径
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	mx.HandleFunc("/", loginHandler(formatter))

	//静态文件服务
	mx.PathPrefix("/static/files").Handler(http.StripPrefix("/static/files", http.FileServer(http.Dir(webRoot+"/assets/"))))
}
