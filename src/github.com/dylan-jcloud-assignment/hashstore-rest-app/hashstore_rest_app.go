package main

import (
	"net/http"
	"log"
	"flag"
	"github.com/dylan-jcloud-assignment/hashstore"
)

var store hashstore.HashStore
var srv http.Server

func main() {
	port := flag.String("port", "8080", "HTTP server port")
	flag.Parse()
	*port = ":" + *port

	store = new(hashstore.SimpleKVHashStore)
	store.Start()

	srv = http.Server{Addr: *port}

	http.HandleFunc("/hash", hashPost)
	http.HandleFunc("/hash/", hashGet)
	http.HandleFunc("/stats", stats)
	http.HandleFunc("/shutdown", shutdown)

	log.Fatal(srv.ListenAndServe())
}
