{{ define "registration_template" }}
    {{ template "header" }}

    <h2>{{ .Title }}</h2>

    <form method="POST" action="/register">
    	Name: <input type="text" name="name"><br/>
        Username: <input type="text" name="username"><br/>
        Password: <input type="password" name="password"><br/><br/>
        <input type="submit" value="Register">
    </form>

    {{ template "footer" }} 
{{ end }}
