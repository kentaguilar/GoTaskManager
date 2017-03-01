package app 

import (
	"net/http"
	"./controller"
)

func Initialize() {
	registration := controller.Registration{ "Sign-up Now" }
	public := controller.Public{ "Main Page" }

	http.HandleFunc("/register", registration.Index)
	http.HandleFunc("/", public.Index)
	http.ListenAndServe(":8080", nil)
}