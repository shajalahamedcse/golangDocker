package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "<h2>Hello World from Distelli! You have a working Go application Deployment!</h2>")
}

func main() {
	port := 8002
	fmt.Println(port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
