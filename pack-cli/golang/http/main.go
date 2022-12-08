package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func html(hostname string) string {
	return `<!DOCTYPE html>
	<html>
	  <head>
		<title>myapp</title>
	  </head>
	  <body>
		<h1>myapp:go:<span style="color:red">v3</span> hostname :　` + hostname + `</h1>
	  </body>
	</html>`
}

func main() {
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, html(hostname))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
