package main

import (
	"fmt"
	"net/http"

	"github.com/nickchou/gnydemo/controllers"
)

func main() {
	var port string = ":10000"

	//http.FileServer(http.Dir("static"))
	http.Handle("/css/", http.FileServer(http.Dir("static")))

	//http.HandleFunc("/static",StaticServer)
	//http.HandleFunc( "/static",StaticServer)
	//注册静态文件
	//http.Handle("/css/", http.FileServer(http.Dir("static")))
	//http.Handle("/script/", http.FileServer(http.Dir("static")))
	http.Handle("/image/", http.FileServer(http.Dir("static")))

	//reg handleFunc,see route.go
	http.HandleFunc("/index", Index)
	http.HandleFunc("/week", controllers.Week)
	http.HandleFunc("/", DefHandle)
	//console
	fmt.Println("listen port" + port)
	//lister port
	http.ListenAndServe(port, nil)
}
func StaticServer(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "text/html")
	staticHandler := http.FileServer(http.Dir("./static/"))
	staticHandler.ServeHTTP(w, r)
	return
}
