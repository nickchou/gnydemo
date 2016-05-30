package main

import (
	"fmt"
	"net/http"
)

func main() {
	var port string = ":10000"

	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/image/", http.FileServer(http.Dir("static")))

	//reg handleFunc,see route.go
	http.HandleFunc("/index", Index)
	http.HandleFunc("/week", Week)
	http.HandleFunc("/", DefHandle)
	//console
	fmt.Println("listen port" + port)
	//lister port
	http.ListenAndServe(port, nil)
}
