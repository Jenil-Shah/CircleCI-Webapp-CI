package main

import (
	"fmt"
	"net/http"
)

func helloWorlds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test CI Webserver 3")
}

func main() {
	http.HandleFunc("/", helloWorlds)
	http.ListenAndServe(":80", nil)
}
