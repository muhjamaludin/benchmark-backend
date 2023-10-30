package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorld struct {
	Hello string
}

func setupRoute() *gin.Engine {
	var message = HelloWorld{
		Hello: "World",
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, message.Hello)
	})

	return r
}

// func handlerIndex(w http.ResponseWriter, r *http.Request) {
// 	var message = HelloWorld{
// 		Hello: "World",
// 	}

// 	b, err := json.Marshal(message)
// 	if err != nil {
// 		fmt.Printf("Error: %s", err)
// 	}

// 	w.Write(b)
// }

func main() {
	// http.HandleFunc("/", handlerIndex)

	// var address = "localhost:3001"
	// err := http.ListenAndServe(address, nil)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// }
	r := setupRoute()
	r.Run(":3001")
}
