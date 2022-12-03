package controller

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var flag bool = false
var files []string

type Client struct {
	id       interface{}
	username string
	email    string
	role     string
	balance  string
}

func (s *DataBase) HomePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap := map[string]any{}
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

	//value_cookie := c.Value
	user, status := s.GetUser(rw, r)
	conditionsMap["username"] = user.username
	conditionsMap["email"] = user.email
	conditionsMap["role"] = user.role
	conditionsMap["balance"] = user.balance
	if status == true {
		conditionsMap["LoginFlagAccept"] = true
	} else {
		conditionsMap["LoginFlagAccept"] = false
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	files = []string{
		"./static/html/home.page.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func (s *DataBase) GetUser(rw http.ResponseWriter, r *http.Request) (Client, bool) {
	session, _ := loggedUserSession.Get(r, "authenticated-user-session")
	userID, ok := session.Values["userID"]
	var user Client
	if ok {

		row := s.Data.QueryRow(`SELECT username, email, role, balance FROM users_account WHERE id = ?`, userID)
		var username, email, role, balance string
		err := row.Scan(&username, &email, &role, &balance)
		_ = err

		user = Client{id: userID, username: username, email: email, role: role, balance: balance}
		return user, true
	}
	return user, false
}

func SignPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap := map[string]any{}
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

type GameHistory struct {
	Tupe       string
	Bet_amount string
	Stat       string
	Summ       string
}

func (s *DataBase) ProfilePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap := map[string]any{}
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

	user, status := s.GetUser(rw, r)
	conditionsMap["username"] = user.username
	conditionsMap["email"] = user.email
	conditionsMap["role"] = user.role
	conditionsMap["balance"] = user.balance

	conditionsMap["history"] = s.getBetHistory(r)

	if status == true {
		conditionsMap["LoginFlagAccept"] = true
	} else {
		conditionsMap["LoginFlagAccept"] = false
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	files = []string{
		"./static/html/profile.page.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}

func (s *DataBase) getBetHistory(r *http.Request) []GameHistory {
	employee := GameHistory{}
	employees := []GameHistory{}

	session, _ := loggedUserSession.Get(r, "authenticated-user-session")
	userID, ok := session.Values["userID"]
	if ok {
		rows, err := s.Data.Query(`SELECT type, bet_amount, stat, summ FROM transactions WHERE user_id=? ORDER BY id DESC`, userID)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var tupe, bet_amount, stat, summ string
			err = rows.Scan(&tupe, &bet_amount, &stat, &summ)
			if err != nil {
				panic(err)
			}
			employee.Tupe = tupe
			employee.Bet_amount = bet_amount
			employee.Stat = stat
			employee.Summ = summ
			employees = append(employees, employee)
		}
	}
	return employees
}

func (s *DataBase) DicePage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap := map[string]any{}
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

	user, status := s.GetUser(rw, r)
	conditionsMap["username"] = user.username
	conditionsMap["email"] = user.email
	conditionsMap["role"] = user.role
	conditionsMap["balance"] = user.balance
	if status == true {
		conditionsMap["LoginFlagAccept"] = true
	} else {
		conditionsMap["LoginFlagAccept"] = false
		http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
		return
	}

	files = []string{
		"./static/html/dice.tmpl",
		"./static/html/basic.layout.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, conditionsMap)
}
