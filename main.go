package main

import (
	"TestGO/controller"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

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

	r.GET("/", controller.HomePage)
	r.GET("/newpage", controller.SecondPage)
	r.GET("/logout", controller.LogoutHandler)
	r.GET("/profile", controller.ProfilePage)
	r.GET("/roulette", controller.RoulettePage)
	r.GET("/dice", controller.DicePage)
	r.GET("/crash", controller.DicePage)
	r.GET("/jackpot", controller.DicePage)
}
