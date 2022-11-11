package controller

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var flag bool = false
var files []string

func HomePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Path != "/" {
		http.NotFound(rw, r)
		return
	}
	if flag == false {
		files = []string{
			"./static/html/home.page.tmpl",
			"./static/html/basic.layout.tmpl",
		}
	} else {
		files = []string{
			"./static/html/home.page.signed.tmpl",
			"./static/html/basic.layout.tmpl",
		}
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, nil)
}

func SecondPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Path != "/newpage" {
		http.NotFound(rw, r)
		return
	}

	if flag == false {
		files = []string{
			"./static/html/newpage.tmpl",
			"./static/html/basic.layout.tmpl",
		}
	} else {
		files = []string{
			"./static/html/newpage.signed.tmpl",
			"./static/html/basic.layout.tmpl",
		}
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, nil)
}

func SignPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Path != "/main-sign" {
		http.NotFound(rw, r)
		return
	}

	var files = []string{
		"./static/html/main_sign.html",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, nil)
}
