package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Benchmark struct {
	Runner        string `json:"runner"`
	Concurrency   int    `json:"concurrency"`
	Max_requests  int    `json:"max_requests"`
	Total_errors  int    `json:"total_erros"`
	Total_time    string `json:"total_time"`
	Mean_latency  string `json:"mean_latency"`
	Effective_rps string `json:"effective_rps"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello World!")
}

func main() {

	// open json
	jsonFile, err := os.Open("benchmark.json")
	// if we os.Open return an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened benchmark.json")

	defer jsonFile.Close()

	// read our opened jsonFile as a byte array
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Users array
	var benchmarks []Benchmark

	// we unmarshal our byteArray
	json.Unmarshal(byteValue, &benchmarks)
	fmt.Println(benchmarks)

	for i := 0; i < len(benchmarks); i++ {
		// total_time, err := strconv.ParseFloat(benchmarks[i].Total_time, 32)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		fmt.Println("Runner: " + benchmarks[i].Runner)
		fmt.Println("Concurrency: " + strconv.Itoa(benchmarks[i].Concurrency))
		fmt.Println("Max Request: " + strconv.Itoa(benchmarks[i].Max_requests))
		fmt.Println("Total errors: " + strconv.Itoa(benchmarks[i].Total_errors))
		fmt.Println("Total time: " + benchmarks[i].Total_time)
		fmt.Println("Mean latency: " + benchmarks[i].Mean_latency)
		fmt.Println("Effective rps: " + benchmarks[i].Effective_rps)
		fmt.Println()
	}
	fmt.Println(len(benchmarks))

	// func main
	// http.HandleFunc("/", hello)
	// fmt.Println("Oke")

	// http.ListenAndServe(":4000", nil)
}
