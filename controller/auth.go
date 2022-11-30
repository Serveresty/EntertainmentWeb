package controller

import (
	"database/sql"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

var encryptionKey = "something-very-secret"
var loggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))

type DataBase struct {
	Data *sql.DB
}

func (s *DataBase) SignUp(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap := map[string]any{}

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
			INSERT INTO users_account (username, email, password, role, balance) VALUES(?, ?, ?, ?, ?)
		`

			//Добавляю в БД запись о регистрации, если нет ошибок
			insert, errdb := s.Data.Query(perem, username, email, hash_password, role, 0.00)
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
			conditionsMap["LoginFlagAccept"] = true
			http.Redirect(rw, r, "/main-sign", http.StatusFound)
		}
	}
}

func (s *DataBase) SignIn(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap := map[string]any{}

	username := r.FormValue("username")
	password := r.FormValue("password")

	conditionsMap["AccessError"] = false
	conditionsMap["WrongPassword"] = false

	row := s.Data.QueryRow(`SELECT id, email, role, balance, password FROM users_account WHERE username = ?`, username)
	if row.Err() != nil {
		rw.Write([]byte("first"))
		rw.Write([]byte(row.Err().Error()))
		return
	}
	var userID, hash, email, role, balance string
	if err := row.Scan(&userID, &email, &role, &balance, &hash); err != nil {
		if err == sql.ErrNoRows {
			conditionsMap["AccessError"] = true
			conditionsMap["LoginFlagAccept"] = false
			http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
			return
		}
		rw.Write([]byte(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			conditionsMap["WrongPassword"] = true
			conditionsMap["LoginFlagAccept"] = false
			http.Redirect(rw, r, "/main-sign", http.StatusSeeOther)
			return
		}
		rw.Write([]byte(err.Error()))
		return
	} else {
		session, _ := loggedUserSession.Get(r, "authenticated-user-session")
		session.Values["userID"] = userID
		session.Save(r, rw)
	}

	conditionsMap["LoginError"] = false
	conditionsMap["LoginFlagAccept"] = true
	http.Redirect(rw, r, "/", http.StatusFound)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *DataBase) LogoutHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session, _ := loggedUserSession.Get(r, "authenticated-user-session")
	delete(session.Values, "userID")
	session.Save(r, rw)
	http.Redirect(rw, r, "/main-sign", http.StatusFound)
}
