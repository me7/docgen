<!doctype html>
<title>Test Document Server</title>
<body>
	<h1>Test Document Center</h1>
	Last update on xxxxx
	<table border="1">
	<tr>
		<th rowspan="2">Product</th>
		<th colspan="3">BOM</th>
		<th colspan="3">Gerber</th>
		<th colspan="3">Schematic</th>
	</tr>
	<tr>
		<th >Group</th>
		<th >REV</th>
		<th >DATE</th>
		<th >Group</th>
		<th >REV</th>
		<th >DATE</th>
		<th >Group</th>
		<th >REV</th>
		<th >DATE</th>
	</tr>
	{{range .}}
	<tr>
		<td>{{.Name}}</td>
		{{range .Docs}}
			<td><a href={{.Url}}>{{.Group}}</a></td>
			<td>{{.Rev}}</td>
			<td>{{.Date}}</td>
			{{end}}
		</td>
	</tr>
	{{end}}
	</table>
</body>