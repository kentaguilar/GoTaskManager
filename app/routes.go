package app 

import (
	"net/http"
	"./controller"
)

func Initialize() {
	public := controller.Public{ "Main Page" }
	task := controller.Task{ "Task Page" }

	http.HandleFunc("/task", task.ModifyOrDeleteTask)
	http.HandleFunc("/", public.Index)
	http.ListenAndServe(":8080", nil)
}