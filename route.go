package main

import (
	"io"
	"log"
	"net/http"
	"html/template"
)
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
func Week(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hi  i'm week")
}