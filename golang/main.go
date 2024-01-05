package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello World!")
}

func main() {

	// func main
	// http.HandleFunc("/", hello)
	// fmt.Println("Oke")

	// http.ListenAndServe(":4000", nil)
}
