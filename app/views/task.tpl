{{ define "task_template" }}
    {{ template "header" }}

    <div style="width:600px;margin:0 auto;text-align:center">
    	<h2>{{ .Title }}</h2> <br/>
    	<input type="text" id="txttask" class="form-control" />
    </div>

    {{ template "footer" }} 
{{ end }}
