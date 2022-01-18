package main

import (
	"github.com/goombaio/namegenerator"
	"net/http"
)

func main() {

	nameGenerator := namegenerator.NewNameGenerator(45)

	http.HandleFunc("/api/qotd", func(rw http.ResponseWriter, r *http.Request) {
		name := nameGenerator.Generate()
		rw.Write([]byte(name))
	})

	http.ListenAndServe(":5000", nil)	
}
