package main

import (
	"io"
	"net/http"
	"strings"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 2 {
		dogName = fs[1]
	}
	// the image doesn't serve
	io.WriteString(res, `
	<h1>Dog Name: `+dogName+`</h1><br>
	<img src="/assets/toby.jpg">
	`)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", upTown)
	http.ListenAndServe(":9000", nil)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
