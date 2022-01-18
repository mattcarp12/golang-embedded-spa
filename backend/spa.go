//go:build spa
// +build spa

package main

import "net/http"

func spaHandler(w http.ResponseWriter, r *http.Request) {

}

func init() {
	http.HandleFunc("/", spaHandler)
}
