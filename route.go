package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("Remote_addr")
	if ip == "" {
		ip = "Addr:" + r.RemoteAddr
	}
	io.WriteString(w, "hello!"+ip)
}
func DefHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index", http.StatusFound)
	}
	t, err := template.ParseFiles("tmpl/error/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
