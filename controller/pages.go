package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var flag bool = false
var files []string

func (s *DataBase) HomePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c, err := r.Cookie("authenticated-user-session")
	if err != nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}
	if c == nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(rw, r)
		return
	}

	value_cookie := c.Value
	ins_sess := `
		INSERT INTO sessions (session_id, user_id, created_at) VALUES(?, ?, ?)
	`
	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	insert, errdb := s.Data.Query(ins_sess, value_cookie, current_client.id, dt)
	defer func() {
		if insert != nil {
		}
	}()
	if errdb != nil {
	}

	files = []string{
		"./static/html/home.page.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	fmt.Println(current_client)
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func SecondPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c, err := r.Cookie("authenticated-user-session")
	if err != nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}
	if c == nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	if r.URL.Path != "/newpage" {
		http.NotFound(rw, r)
		return
	}

	value_cookie := c.Value
	fmt.Println(value_cookie)

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
	c, err := r.Cookie("authenticated-user-session")
	if err != nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}
	if c == nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	if r.URL.Path != "/profile" {
		http.NotFound(rw, r)
		return
	}

	value_cookie := c.Value
	fmt.Println(value_cookie)

	files = []string{
		"./static/html/profile.page.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func RoulettePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c, err := r.Cookie("authenticated-user-session")
	if err != nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}
	if c == nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	if r.URL.Path != "/roulette" {
		http.NotFound(rw, r)
		return
	}

	value_cookie := c.Value
	fmt.Println(value_cookie)

	files = []string{
		"./static/html/newpage.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func DicePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c, err := r.Cookie("authenticated-user-session")
	if err != nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}
	if c == nil {
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	if r.URL.Path != "/dice" {
		http.NotFound(rw, r)
		return
	}

	value_cookie := c.Value
	fmt.Println(value_cookie)

	files = []string{
		"./static/html/dice.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}
