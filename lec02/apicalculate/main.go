package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Object struct {
	Number1     float64 `json:"number1"`
	Number2     float64 `json:"number2"`
	Expresstion string  `json:"expression"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	var o Object
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, "Err", http.StatusBadRequest)
		return
	}
	var result float64
	switch o.Expresstion {
	case "+":
		result = o.Number1 + o.Number2
	case "-":
		result = o.Number1 - o.Number2
	case "*":
		result = o.Number1 * o.Number2
	case "/":
		result = o.Number1 / o.Number2
	default:
		result = 0
	}
	fmt.Fprintf(w, "%v", result)
	
}

func main() {
	http.HandleFunc("/cal", Calculate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}