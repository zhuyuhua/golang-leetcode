package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"hello," ,html.EscapeString(request.URL.Path))
	})

	log.Fatal(http.ListenAndServe("localhost:8080",nil))
	
}
