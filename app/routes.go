package app 

import (
	"net/http"
	"./controller"
)

// var db *sql.DB
// var err error

func Initialize() {
	// db, err = sql.Open("mysql", "root:@/godb")
 //    if err != nil{
 //        panic(err.Error())
 //    }
 //    defer db.Close()

 //    err = db.Ping()
 //    if err != nil{
 //        panic(err.Error())
 //    }

	registration := controller.Registration{ "Sign-up Now" }
	public := controller.Public{ "Main Page" }

	http.HandleFunc("/register", registration.Index)
	http.HandleFunc("/", public.Index)
	http.ListenAndServe(":8080", nil)
}