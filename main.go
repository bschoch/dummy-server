package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dat, _ := ioutil.ReadFile("./response.json")
		fmt.Fprint(w, string(dat))
		os.Rename("./response.json", "./response.json.bak")
	})

	http.ListenAndServe(":8080", mux)
}
