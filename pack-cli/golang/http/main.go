package main

import (
	"fmt"
	"log"
	"net/http"
)

const INDEX = `<!DOCTYPE html>
<html>
  <head>
    <title>myapp</title>
  </head>
  <body>
    <h1>myapp:go:<span style="color:red">v3</span></h1>
  </body>
</html>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, INDEX)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
