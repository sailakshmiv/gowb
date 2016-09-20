package main

import (
	"net/http"
	"text/template"
)

func main()  {
	http.HandleFunc("/",func(w http.ResponseWriter,req *http.Request){
		w.Header().Add("Cotent type", "text/html")
		tmpl,err:= template.New("test").Parse(doc)
		if err == nil{
			context:= Context{Message:"Sample web App",Title:"Title of Hello world"}
			tmpl.Execute(w,context)
		}
	})
	http.ListenAndServe(":8080",nil)
}

const doc = `
<!DOCTYPE html>
<html>
<head>
<title> {{.Title}}</title>
</head>
<body>
<h1> {{.Message}}</h1>
</body>
</html>
`

type Context struct {
	Title string
	Message string
}
