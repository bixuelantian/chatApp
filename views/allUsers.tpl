<head>

</head>

<body>
	<table align="center" border="1">
		<tr>
			<th>Id</th>
			<th>Username</th>
			<th>Email</th>
		</tr>

		{{range .users}}
			<tr>
				<td>{{.Id}}</td>
				<td>{{.Username}}</td>
				<td>{{.Email}}</td>
			</tr>
		{{end}}

	<table>

	<div align="center">
	</div>
<body>
