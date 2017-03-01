package controller

import (
	// "fmt"
	"net/http"
	"../helpers"
	"../models"	
)

type Registration struct {
	Title string
}

func (registration Registration) Index(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		view := helper.RenderPage(response, "registration.tpl")

		if err := view.ExecuteTemplate(response, "registration_template", registration); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}

    	return
	}

	fullname := request.FormValue("name")
	username := request.FormValue("username")
    password := request.FormValue("password")

    user := model.User{ fullname, username, password }
    output := user.SaveUser()

    response.Write([]byte(output))
}