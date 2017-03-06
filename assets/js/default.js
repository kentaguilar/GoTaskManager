$(function(){
	$('#txttask').on('keypress', function(e){
		if(e.which == 13) {
			if($('#txttask').val().length > 0){
				if($('#txttask').attr('taskid') == "" || $('#txttask').attr('taskid') == undefined){
					modifyTask("new");					
				}
				else{
					modifyTask("update");
				}

				$('#txttask').attr('taskid','')
			}
		}
	});

	showAllTasks();

	function modifyTask(type){
		$.ajax({
		    url: "crud?name=" + $('#txttask').val() + "&type=" + type + (type == "update" ? "&id=" + $('#txttask').attr('taskid')  : ""),
		    type: "GET",
		    success: function(e){
		    	showAllTasks();
		    	$('#txttask').val('').focus();
		    },
		    error: function(req, status, err) {
		    	console.log("error found", status, err);
		    }
		});
	}

	function showAllTasks(){
		$.ajax({
		    url: "tasks",
		    type: "GET",
		    success: function(e){
		    	if(e.length > 0){
		    		var output = "";
		    		for(var i = 0; i < e.length; i++){
		    			output += "<div taskid='" + e[i].Id + "'><span>" + e[i].Name + "</span><i class='delete pull-right fa fa-close'></i>";
		    			output += "<i class='edit pull-right fa fa-edit'></i></div>";
		    		}

		    		$('.task-list').empty().append(output);

		    		$('.task-list div i.delete').on('click', function(){
		    			currentId = $(this).parent().attr('taskid');

		    			$.ajax({
						    url: "crud?type=delete&id=" + currentId,
						    type: "GET",
						    success: function(e){
						    	showAllTasks();
						    },
						    error: function(req, status, err) {
						    	console.log("error found", status, err);
						    }
						});
		    		});

		    		$('.task-list div i.edit').on('click', function(){
		    			currentText = $(this).parent().find('span').text();
		    			currentId = $(this).parent().attr('taskid');

		    			$('#txttask').val(currentText).attr('taskid', currentId);
		    		});
		    	}
		    	else{
		    		$('.task-list').empty().append("<div>No tasks yet</div>");
		    	}
		    },
		    error: function(req, status, err) {
		    	console.log("error found", status, err);
		    }
		});
	}	
});