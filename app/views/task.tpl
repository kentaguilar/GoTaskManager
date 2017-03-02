{{ define "task_template" }}
    {{ template "header" }}

    <h2>{{ .Title }}</h2>    

    <div>
    	<input type="text" id="txttask" />
    </div>

    {{ template "footer" }} 
{{ end }}
