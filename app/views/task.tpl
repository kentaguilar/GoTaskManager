{{ define "task_template" }}
    {{ template "header" }}

    <div style="width:600px;margin:0 auto;text-align:center">
    	<h2>{{ .Title }}</h2> <br/>
    	<input type="text" id="txttask" class="form-control" placeholder="Add your task here" />

    	<div class="task-list"></div>
    </div>    

    {{ template "footer" }} 
{{ end }}
