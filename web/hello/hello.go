package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	io.WriteString(w, "hello world!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListentAndServer: ", err.Error())
	}
}
