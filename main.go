package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dat, _ := ioutil.ReadFile("./response.json")
		fmt.Fprint(w, string(dat))
	})

	http.ListenAndServe(":8080", mux)
}
