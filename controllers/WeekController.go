package main

import (
	"net/http"
	"io"
)
type WeekController struct {
}

func (this *WeekController) IndexAction(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hi  i'm week")
}