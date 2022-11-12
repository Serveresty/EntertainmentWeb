package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

var encryptionKey = "something-very-secret"
var loggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))

type any interface{}

var conditionsMap map[string]any

type Complete struct {
	Data *sql.DB
}

func init() {

	loggedUserSession.Options = &sessions.Options{
		// change domain to match your machine. Can be localhost
		// IF the Domain name doesn't match, your session will be EMPTY!
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 3, // 3 hours
		HttpOnly: true,
	}
}

func (s *Complete) SignUp(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap = map[string]any{}

	session, _ := loggedUserSession.Get(r, "authenticated-user-session")

	if session != nil {
		conditionsMap["username"] = session.Values["username"]
	}

	//Беру данные при регистрации
	if r.FormValue("email") != "" && r.FormValue("username") != "" && r.FormValue("password") != "" && r.FormValue("confirm_password") != "" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		confirm_password := r.FormValue("confirm_password")
		email := r.FormValue("email")
		role := "user"

		conditionsMap["EmailUsernameError"] = false

		if password != confirm_password {
			conditionsMap["LoginError"] = true
			conditionsMap["LoginFlagAccept"] = false
			http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
			return
		} else {

			//Хэширую пароль
			hash_password, _ := HashPassword(password)

			//Переменная с кодом MySql
			perem := `
			INSERT INTO users_account (username, email, password, role) VALUES(?, ?, ?, ?)
		`

			//Добавляю в БД запись о регистрации, если нет ошибок
			insert, errdb := s.Data.Query(perem, username, email, hash_password, role)
			defer func() {
				if insert != nil {
					insert.Close()
				}
			}()
			if errdb != nil {
				if strings.Contains(errdb.Error(), "Error 1062") {
					conditionsMap["EmailUsernameError"] = true
					conditionsMap["LoginFlagAccept"] = false
					http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
					return
				}
				rw.Write([]byte(errdb.Error())) //Поменять
				return
			}

			conditionsMap["LoginError"] = false
			conditionsMap["username"] = username

			session, _ := loggedUserSession.New(r, "authenticated-user-session")
			session.Values["username"] = username
			err := session.Save(r, rw)
			if err != nil {
				rw.Write([]byte("GWWWW"))
			}

			conditionsMap["LoginFlagAccept"] = true
			http.Redirect(rw, r, "/", http.StatusFound)
		}
	}
}

func (s *Complete) SignIn(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap = map[string]any{}

	session, _ := loggedUserSession.Get(r, "authenticated-user-session")

	if session != nil {
		conditionsMap["username"] = session.Values["username"]
	}

	email := r.FormValue("username")
	password := r.FormValue("password")

	conditionsMap["AccessError"] = false
	conditionsMap["WrongPassword"] = false

	row := s.Data.QueryRow(`SELECT password FROM users_account WHERE email = ?`, email)
	if row.Err() != nil {
		rw.Write([]byte("first"))
		rw.Write([]byte(row.Err().Error()))
		return
	}
	var result string
	if err := row.Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			conditionsMap["AccessError"] = true
			conditionsMap["LoginFlagAccept"] = false
			http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
			return
		}
		rw.Write([]byte(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			conditionsMap["WrongPassword"] = true
			conditionsMap["LoginFlagAccept"] = false
			http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
			return
		}
		rw.Write([]byte(err.Error()))
		return
	}

	conditionsMap["LoginError"] = false
	conditionsMap["username"] = email

	session, _ = loggedUserSession.New(r, "authenticated-user-session")
	session.Values["username"] = email
	err := session.Save(r, rw)
	if err != nil {
		rw.Write([]byte("GWWWW"))
	}
	conditionsMap["LoginFlagAccept"] = true
	http.Redirect(rw, r, "/", http.StatusFound)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func LogoutHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//read from session
	session, _ := loggedUserSession.Get(r, "authenticated-user-session")
	for k := range conditionsMap {
		delete(conditionsMap, k)
	}

	err := session.Save(r, rw)

	if err != nil {
		log.Println(err)
	}
	conditionsMap["LoginFlagAccept"] = false
	http.Redirect(rw, r, "/", http.StatusFound)
}
