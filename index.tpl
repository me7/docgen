<!doctype html>
<title>Test Document Server</title>
<body>
	<h1>Test Document Center</h1>
	Last update on {{.LastUpdate}} <a href="history.txt">History of Changes</a>
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
	{{range $key, $value := .Doc}}
	<tr>
		<td>{{$key.Customer}}</td>
		<td>{{$key.Product}}</td>
		<td>{{$key.DocType}}</td>
		<td>{{$value}}</td>
		</td>
	</tr>
	{{end}}
	</table>
</body>