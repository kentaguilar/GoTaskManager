{{ define "login_template" }}
    {{ template "header" }}

    <h1>Login Page</h1>
    <form method="POST" action="/login">
        <input type="text" name="username" placeholder="username">
        <input type="password" name="password" placeholder="password">
        <input type="submit" value="Submit">
    </form>

    {{ template "footer" }} 
{{ end }}
