package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",solvePing)
	http.ListenAndServe(":9999",nil)
}

func solvePing(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world!")
}
