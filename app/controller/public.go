package controller

import (
	// "net/http"
	// "../helpers"
	// "strings"
	// "fmt"
	// "log"
)

// var chttp = http.NewServeMux()

type Public struct {
	Title string
}

// func (public Public) Index(response http.ResponseWriter, request *http.Request) {
// 	view := helper.RenderPage(response, "main.tpl")

// 	if (strings.Contains(request.URL.Path, ".")) {
//         chttp.ServeHTTP(response, request)
//         //fmt.Println("asdf")
//     } else {
// 		if err := view.ExecuteTemplate(response, "main_template", public); err != nil {
// 			http.Error(response, err.Error(), http.StatusInternalServerError)
// 		}
// 	}

// 	// chttp.ServeHTTP(response, request)
// }

// func Index(h http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     log.Println("Before")
//     h.ServeHTTP(w, r) // call original
//     log.Println("After")
//   })
// }