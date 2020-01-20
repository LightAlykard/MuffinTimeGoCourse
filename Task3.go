package main

import (
    "io"
    "net/http"
    //"fmt"
)

func hello(res http.ResponseWriter, req *http.Request) {
    name := req.URL.Query().Get("name")
	if name == "" {
		name = "Default"
	}
    res.Header().Set("Content-Type", "text/html")
    io.WriteString(res,
        `<doctype html>
<html>
    <head>
        <title>Hello </title>
    </head>
    <body>
        Hello `+name+`! 
    </body>
</html>`)    
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)
    http.HandleFunc("/hello", hello)

    http.ListenAndServe(":80", nil)
}