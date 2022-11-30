package main

import (
	"TestGO/controller"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

var encryptionKey = "something-very-secret"
var loggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))

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

func main() {
	DB, errdb := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users_data")
	if errdb != nil {
		panic(errdb)
	}
	defer DB.Close()
	if err := DB.Ping(); err != nil {
		panic(err)
	}

	r := httprouter.New()
	routes(r, DB)

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router, DB *sql.DB) {
	r.ServeFiles("/static/*filepath", http.Dir("static"))

	handler := controller.DataBase{Data: DB}
	//test SignIn, SignOut
	r.GET("/main-sign", controller.SignPage)
	r.POST("/main-sign-up", handler.SignUp)
	r.POST("/main-sign-in", handler.SignIn)
	//

	r.GET("/", handler.HomePage)
	r.GET("/logout", handler.LogoutHandler)
	r.GET("/profile", handler.ProfilePage)
	r.GET("/dice", handler.DicePage)

	r.POST("/dice", handler.GetDiceData)
	r.POST("/profile", handler.Deposit)
}
