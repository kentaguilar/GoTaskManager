package controller

import (
	"net/http"
	"../helpers"
)

type Public struct {
	Title string
}

func (public Public) Index(response http.ResponseWriter, request *http.Request) {
	view := helper.RenderPage(response, "main.tpl")

	if err := view.ExecuteTemplate(response, "main_template", public); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}