{{ define "main_template" }}
    {{ template "header" }}

    <h2>{{ .Title }}</h2>
    <a href="/login">Login</a><br/>
    <a href="/register">Sign Up</a>

    {{ template "footer" }} 
{{ end }}
