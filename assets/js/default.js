$(function(){
	$('#txttask').on('keypress', function(e){
		if(e.which == 13) {
			console.log("asdf");
		}
	});

	$.ajax({
	    url: "tasks",
	    type: "GET",
	    success: function(e){
	    	if(e.length > 0){
	    		var output = "";
	    		for(var i = 0; i < e.length; i++){
	    			output += "<div id='" + e[i].Id + "'>" + e[i].Name + "</div>";
	    		}

	    		$('.task-list').empty().append(output);
	    	}
	    }
	});
});