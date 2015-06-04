<!doctype html>
<title>Test Document Server</title>
<style>
h2{text-align:center;}
p{text-align: right}
body{
    background: linear-gradient(#CEE2F7, white); /* Standard syntax (must be last) */
}
</style>
<body>
	<h2>Test Document Center</h2>
	<p>Last update on {{.LastUpdate}} <a href="history.txt">History of Changes</a></p>
	<table border="1">
	<tr>
		<th>Customer</th>
		<th>Product</th>
		<th>DocType</th>
		<th>Rev</th>
		<th>Document Name</th>
		<th>Archive</th>
		<th>Register Date</th>
	</tr>	
	{{range $key, $val := .Doc}}
	<tr>
		<td>{{$key.Customer}}</td>
		<td>{{$key.Product}}</td>
		<td>{{$key.DocType}}</td>
		<td>{{$val.Rev}}</td>
		<td><a href="{{$key.Customer}}/{{$key.Product}}/{{$key.DocType}}/{{$val.Rev}}/{{$val.Name}}">{{$val.Name}}</a></td>
		<td><a href="{{$key.Customer}}/{{$key.Product}}/{{$key.DocType}}/">Old Version</a></td>
		<td>{{$val.Date}}</td>
	</tr>	
	{{end}}
	</table>
</body>