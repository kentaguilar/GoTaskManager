package app 

import (
	"net/http"
	"./controller"
	"strings"
)

var chttp = http.NewServeMux()
var task = controller.Task{"Manage My Tasks"}

func Initialize() {
	chttp.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/tasks", task.GetAllTasks)
	http.HandleFunc("/", HomePageHandler)
	http.ListenAndServe(":8080", nil)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
    if (strings.Contains(r.URL.Path, ".")) {
        chttp.ServeHTTP(w, r)
    } else {
        task.Index(w, r)
    }   
} 