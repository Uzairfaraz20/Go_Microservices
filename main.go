package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "error!", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("error!"))
			return
		}

		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye  world")
	})

	http.ListenAndServe(":9090", nil)
}
