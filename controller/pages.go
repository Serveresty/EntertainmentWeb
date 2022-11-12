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
	files = []string{
		"./static/html/home.page.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func SecondPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Path != "/newpage" {
		http.NotFound(rw, r)
		return
	}
	files = []string{
		"./static/html/newpage.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func SignPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Path != "/main-sign" {
		http.NotFound(rw, r)
		return
	}

	var files = []string{
		"./static/html/main_sign.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func ProfilePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Path != "/profile" {
		http.NotFound(rw, r)
		return
	}
	files = []string{
		"./static/html/profile.page.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}
