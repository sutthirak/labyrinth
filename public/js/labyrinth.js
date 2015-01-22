var LABYRINTH = LABYRINTH || {
	ELEMENT:{
		navbar : function(){
			var path = window.location.pathname;
			if(path.substring(path.length-1) === "/"){
				path = path.substring(0,path.length-1);
			}
			var query = 'nav ul li a';
			$(query).each(function(){
				if($(this).attr("href") === path.substring(0,$(this).attr("href").length)){
					$(this).parent().addClass("active");		
				}
			});
		},
		tab : function(){
			var found = false;
			if(window.location.href.indexOf("?") != -1){
				var query = window.location.href.split("?")[1].split("&")
				query.forEach(function(element){
					token = element.split("=");
					if(token[0] === "tab"){
						$("#mainTab li a").each(function(){
							if($(this).text().toLowerCase() === token[1].toLowerCase()){
								$(this).parent().addClass("active");
								$(".tab-pane").each(function(index,panel){
									if(this.id == token[1]+"Panel"){
										$(this).addClass("active");
									}
								});
								found = true;
							}
						}); 
					}
				});	
			}
			if(!found){
				$("#mainTab li a[href='#objectPanel'").parent().addClass("active");
				$("#objectPanel").addClass("active");
			}	
		},
		slider : function(id,startValue){
			var element = document.querySelector(id);
			var doc = {min:0,max:1,step:0.01,start:startValue,hideRange:true,decimal:true};
			var init = new Powerange(element,doc);
		},
		enableField : function(id,menu){
			var process = function(){
				var kind = $(id).val();
				if(kind != ""){
			    	for(var key in menu){
		    			for(var index in menu[key]){
		    				$(menu[key][index]).prop("disabled",true);
		    			}	    	
			    	}

			    	var list = menu[kind];
					for (var index in list){
			    		$(list[index]).prop("disabled",false);
					}
				}
			};
			process();
			$(id).on("change", process);
		}
	},
	MESSAGE: {
		success : function(message){
			new PNotify({title: 'Success',text: message,type: 'info'});
		},
		error : function(message){
			$("#message").append("<button type=\"button\" class=\"close\" data-dismiss=\"alert\"><span aria-hidden=\"true\">&times;</span>");
			$('#message').addClass("alert alert-danger alert-dismissible").attr("role","alert");
			$('#message').html("<strong>"+message+"</strong>");
		}	
	},
	ACTION:{
		post : function(endpoint){
			$.ajax({
			    url: endpoint,
			    type: "POST",
			    success: function() {
			    	location.reload(true);
			    },
				error: function(result){
					LABYRINTH.MESSAGE.error(result);
				}
			});
		},
		send : function(endpoint,body){
			$.ajax({
			    url: endpoint,
			    type: "POST",
			    data: body,
          		contentType:"text/plain",
			    success: function() {
			    	LABYRINTH.MESSAGE.success("Done!");
			    },
				error: function(result){
					LABYRINTH.MESSAGE.error(result);
				}
			});
		},
		delete : function(endpoint){
			if(confirm("Do you want to delete ?")) {
				$.ajax({
				    url: endpoint,
				    type: "DELETE",
				    success: function() {
				    	location.reload(true);
				    },
					error: function(result){
						LABYRINTH.MESSAGE.error(result);	
					}
				});
			}
		}
	}

};