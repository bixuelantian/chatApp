<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>pleas login</title>
</head>
<body>
  <div style="text-align:center;font-size:16px;padding-top:50px;padding-bottom:50px;">
	This is we chat <br/>
	Please fill your register info
 </div>
  <form id="login-form" method="post" action="/register" align="center">
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
	<tr>
	   <td>
    		<label>email:</label>
	   </td>
	   <td>
    		<input name="email"></input>
	   </td>
	<tr>
	<tr>
	   <td colspan="2">
    		<button type="submit">register in</button>
	   </td>
	</tr>
  </form>
</body>
</html>
