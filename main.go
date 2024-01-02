package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var directory, bind string
	flag.StringVar(&directory, "directory", "./", "root directory")
	flag.StringVar(&bind, "bind", "127.0.0.1:8080", "bind address")
	flag.Parse()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(directory))))
	log.Println("Listen on", bind, "server directory", directory)
	log.Fatal(http.ListenAndServe(bind, nil))
}
