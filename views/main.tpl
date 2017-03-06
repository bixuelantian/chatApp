<head>
 	<link rel="stylesheet" href="//code.jquery.com/ui/1.10.4/themes/smoothness/jquery-ui.css">
	<script src="//code.jquery.com/jquery-1.9.1.js"></script>
	<script src="//code.jquery.com/ui/1.10.4/jquery-ui.js"></script>
	<script>

		function add_friend(cf){
			$("#dv_add_friend").dialog(
			 {  
			    width:400,
			    height:200,
		            modal:true,
			    closeOnEscape:false,
			    buttons:{
                    		"Submit":function(){
					var fname = $("#add_fname").val();
					if(fname == null || fname == ""){
						alert("Username is empty");
						return false;
					}					
					var uid = $("#dv_add_friend").attr("userid");
					$.post("/addfriend", {"userid":uid,"friendname":fname},function(d){ if(d.retcode == 0) { alert("add success");}else{ alert(d.retmsg);} }, "json");
				},
                    		"Cancel":function(){$(this).dialog('close');window.location.reload();},
			    }
			 }
			);
			return false;
		}
		function del_friend(cur_f){
			var tr = $(cur_f).closest("tr");	
		        var rid = tr.attr('rid');
			$.post("/delfriend", {"rid":rid}, function(d){ if(d.retcode == 0) {alert("delete friend success");window.location.reload(); }else{alert("delete friend failed")} });	
			return false;
		}

		function chat(cur_f){	
			var tr = $(cur_f).closest("tr");	
		        var fid = tr.attr('fid');
			var tt = $(cur_f).closest("table");
		        var uid = tt.attr('uid');
			
			var msg_str = "";
			$.post("/getmsg", {"uid":uid, "fid": fid}, function(d){ 
				var msgs = new Array();
				var m;
				for(var key in d.message){
					var m="";
					m += d.message[key]["Sendtime"] + " (Id: ";
					m += d.message[key]["Sender"] + ")\n";
					m += d.message[key]["Message"] + "\n";
					msgs.push(m);
				}
				msgs.sort();
				for(var i = 0; i < msgs.length; i++){
					msg_str += msgs[i] + "\n";
				}
				$("#msg").text(msg_str);
				$("#dv_chat").attr("fid", fid);
				$("#dv_chat").attr("title", "Wechat ( " + uid + " <-> " + fid + " )");
				$("#dv_chat").dialog(
        	                 {  
                	            width:500,
                        	    height:420,
				    modal:true,
                          	    buttons:{
                                	"Cancle":function(){$(this).dialog('close');window.location.reload();}
                            	    }
                                 }
                           );
			});
			return false;

		}
		function send_msg(cur){
			var mm = $("#input_msg").val();
			if(mm == null || mm == ""){
				alert("input is empty");
				return;
			}
			$("#input_msg").val("");
			
			var dd = $(cur).closest("div");        
                        var fid = dd.attr('fid');
                        var uid = dd.attr('uid');
			$.post("/sendmsg", {"uid":uid, "fid":fid, "msg":mm, "cur_time":getNow()});
			var  msg_str = $("#msg").val();
			msg_str += getNow() +"(Id: " + uid + ")\n" + mm + "\n\n";
			$("#msg").text(msg_str);
			
			return false;
		}

		function getNow() {
 		   	var date = new Date();
    			var seperator1 = "-";
		    	var seperator2 = ":";
    			var month = date.getMonth() + 1;
		        var strDate = date.getDate();
    			if (month >= 1 && month <= 9) {
        			month = "0" + month;
    			}
    			if (strDate >= 0 && strDate <= 9) {
        			strDate = "0" + strDate;
    			}
    			var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate + " " + date.getHours() + seperator2 + date.getMinutes() + seperator2 + date.getSeconds();
    			return currentdate;
		}
	</script>
	
</head>
<body>
	<div style="padding-left:180px;align:left">
		(userid: {{.user.Id}} &nbsp;&nbsp;username: {{.user.Username}} &nbsp;&nbsp;Email: {{.user.Email}})&nbsp;&nbsp;&nbsp;<a href="/logout">Logout</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	</div>
     	<div style="padding-left:200px;align:left">
             <h1>Welcome to your home!</h1>
     	</div>
    	<div style="padding-left:100px">
        	<table  uid={{.user.Id}} align="left" border="1px" width="600" cellspacing="0" cellpadding="0">
			<caption><font size=6>Friends List</font> &nbsp;&nbsp;&nbsp;(<a href="" onclick="return add_friend(this)">Add Friend</a>)</caption>
                	<tr>
				<th>DEL</th>
                        	<th>Id</th>
                        	<th>Username</th>
                        	<th>Email</th>
				<th>UReadMsg</th>	
				<th>Chat</th>
                	</tr>
			
                	{{range $key, $value := .friends}}
                         <tr style="text-align:center" rid={{$value.Rid}} fid={{$value.Id}}>
				<td><a href="" onclick="return del_friend(this)">DEL</a></td>
                                <td>{{$value.Id}}</td>
                                <td>{{$value.Username}}</td>
                                <td>{{$value.Email}}</td>
				<td>{{$value.Msgnum}}</td>
				<td><a href="" onclick="return chat(this)">Chat</a></td>
                         </tr>	
               	 	{{end}}

        	</table>
    	</div>

	<div id="dv_add_friend"  title="Add Friend" hidden="hidden" userid={{.user.Id}}>
		<label>Username:</label>
		<input id="add_fname"></input>
	</div>

	<div id="dv_chat"  title="WeChat" hidden="hidden" uid={{.user.Id}} fid="">
		<table>
		<tr>
			<td colspan="2"><textarea readonly="readonly" id="msg"  op_type="" cols="45" rows="12"></textarea></td>
		</tr>
		<tr>
			<td><input id="input_msg"></input></td><td><button onclick="return send_msg(this)">submit</button></td>
		<tr>
		</table>
        </div>
</body>
