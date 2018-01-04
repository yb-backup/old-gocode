package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port to server on")
	dir := flag.String("d", ".", "the static dir")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Println("listen 0.0.0.0:", *port, " dir:", *dir)
	http.ListenAndServe("0.0.0.0:"+*port, nil)
}
