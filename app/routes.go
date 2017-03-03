package app 

import (
	"net/http"
	"./controller"
)

func Initialize() {
	public := controller.Public{ "Main Page" }
	task := controller.Task{ "Manage My Tasks" }

	http.HandleFunc("/task", task.GetAllTasks)
	http.HandleFunc("/task", task.Index)
	http.HandleFunc("/", public.Index)
	http.ListenAndServe(":8080", nil)
}