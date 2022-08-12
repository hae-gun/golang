package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":3000"
	fmt.Println("Server start at port%v", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		fmt.Fprint(w, "Hello World")
	})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		fmt.Fprint(w, "Hello bar!")
	})
	http.ListenAndServe(port, nil)
}
