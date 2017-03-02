package controller

import (
	"net/http"
	"../helpers"
	"../models"	
	"encoding/json"
)

type Task struct {
	Title string
}

type Profile struct {
  Name    string
  Hobbies []string
}


func (task Task) Index(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		view := helper.RenderPage(response, "task.tpl")

		if err := view.ExecuteTemplate(response, "task_template", task); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}

    	return
	}
}

func (task Task) ModifyOrDeleteTask(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	description := request.FormValue("description")
	id := request.FormValue("id")
	requestType := request.FormValue("type")

	newTask := model.Task{ id, name, description }
	output := ""

	if requestType == "new" {
		output = newTask.SaveTask()
	} else if requestType == "update" {
		output = newTask.UpdateTask()
	} else {
		output = newTask.DeleteTask()
	}

	js, err := json.Marshal(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}