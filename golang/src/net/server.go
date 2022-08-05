package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintf(w, "안녕하세요! %s에 오신 것을 완영합니다!", r.URL.Path[1:])
}
