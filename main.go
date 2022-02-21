package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	//create logger and http handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	//create new servemux
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", nil)
}
