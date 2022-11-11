package controller

import (
	"database/sql"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type Complete struct {
	Data *sql.DB
}

func (s *Complete) SignUp(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Беру данные при регистрации
	username := r.FormValue("username")
	password := r.FormValue("password")
	//confirm_password := r.FormValue("confirm_password")
	email := r.FormValue("email")
	role := "user"

	/*if password != confirm_password {
		http.Redirect(rw, r, "/register", http.StatusSeeOther) //Исправить
		return
	}*/

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
			rw.Write([]byte("ERRRROOOOORRRR")) //Поменять
			return
		}
		rw.Write([]byte(errdb.Error())) //Поменять
		return
	}
	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

func (s *Complete) SignIn(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	row := s.Data.QueryRow(`SELECT password FROM users_account WHERE email = ?`, email)
	if row.Err() != nil {
		rw.Write([]byte("first"))
		rw.Write([]byte(row.Err().Error()))
		return
	}
	var result string
	if err := row.Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			rw.Write([]byte("User not found"))
			return
		}
		rw.Write([]byte(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			rw.Write([]byte("User not found"))
			return
		}
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte("User found")) //Поменять
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
