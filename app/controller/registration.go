package controller

import (
	"net/http"
)

type Registration struct {

}

func (registration Registration) Index(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.ServeFile(response, request, "views/register.html")
		return
	}

	username := request.FormValue("username")
    password := request.FormValue("password")

    fmt.Println(username)
}