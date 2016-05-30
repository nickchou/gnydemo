package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func main() {
	var port string =  ":10000"
	
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/image/", http.FileServer(http.Dir("static")))
	//reg handle
	http.HandleFunc("/index",Index)
	http.HandleFunc("/",DefHandle)
	//console
	fmt.Println("listen port"+port)
	//lister port
	http.ListenAndServe(port, nil)
	
}
func Index(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello  golang!")
}
func DefHandle(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" {
        http.Redirect(w, r, "/index", http.StatusFound)
    }
    t, err := template.ParseFiles("tmpl/error/404.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}
