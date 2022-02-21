package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//struct Hello that contains logger
type Hello struct {
	l *log.Logger
}

//creates and returns Hello object with a logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "error!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "hello %s\n", d)
}
