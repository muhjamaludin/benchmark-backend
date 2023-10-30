package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloWorld struct {
	Hello string
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = HelloWorld{
		Hello: "World",
	}

	b, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	w.Write(b)
}

func main() {
	http.HandleFunc("/", handlerIndex)

	var address = "localhost:3001"
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
