package helper

import (
	"html/template"
	"net/http"
	"path"
)

const(
    BASEURL = "app/views/"
    HEADERTEMPLATE = "app/views/layouts/header.tpl"
    FOOTERTEMPLATE = "app/views/layouts/footer.tpl"
)

func RenderPage(response http.ResponseWriter, targetTemplate string) *template.Template {
	content := path.Join("app", "views", targetTemplate)

	tmpl, err := template.ParseFiles(content, HEADERTEMPLATE, FOOTERTEMPLATE)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	return tmpl	
}