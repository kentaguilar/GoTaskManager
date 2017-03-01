package app 

import (
	"net/http"
	"./controller"
)

func Initialize() {
	registration := controller.Registration {}

	http.HandleFunc("/registration", registration.Index)
	http.ListenAndServe(":8080", nil)
}