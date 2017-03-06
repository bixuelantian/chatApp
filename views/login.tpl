<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>pleas login</title>
</head>
<body>
  <div style="text-align:center;font-size:16px;padding-top:50px;padding-bottom:50px;">
	This is wechat
  <br/>
	Please fill your loginfo
 </div>
  <form id="login-form" method="post" action="/login" align="center">
    <table align="center">
	<tr>
	   <td>
    		<label>Username:</label>
	   </td>
	   <td>
    		<input name="username" ></input>
	   </td>
        </tr>
	<tr>
	   <td>
    		<label>Password:</label>
	   </td>
	   <td>
    		<input type="password" name="password"></input>
	   </td>
	<tr>
	   <td colspan="2">
    		<button type="submit">Login in</button>
	   </td>
	</tr>
  </form>
</body>
</html>
