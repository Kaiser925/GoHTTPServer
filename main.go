package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var directory string
var bind string
var port int

func init() {
	flag.StringVar(&directory, "directory", "./", "root directory")
	flag.StringVar(&bind, "bind", "127.0.0.1", "bind address")
	flag.IntVar(&port, "port", 8765, "bind port")
	flag.Parse()
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(directory))))
	addr := fmt.Sprintf("%s:%d", bind, port)
	log.Println("Listen on", addr, "server directory", directory)
	log.Fatal(http.ListenAndServe(addr, nil))
}
