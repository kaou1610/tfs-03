package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Object struct {
	Number1 float64 `json:"number_1"`
	Number2 float64 `json:"number_2"`
	Expression string `json: "expression"`

}
func handlerPost(w http.ResponseWriter, r *http.Request){
	var o Object
	err := json.NewDecoder(r.Body).Decode(&o)
	if err!= nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	result, err := calculate(&o)
	if err!=nil {
		fmt.Fprintf(w, "%v", err.Error())
		return
	}
	fmt.Fprintf(w, "%v", result)
}

func calculate(o *Object) (float64, error)  {
	switch o.Expression {
	case "+":
		return o.Number1 + o.Number2, nil
	case "-":
		return o.Number1 + o.Number2, nil
	case "*":
		return o.Number1 + o.Number2, nil
	case "/":
		return o.Number1 + o.Number2, nil
	}
	return 0, nil
}

func main()  {
	
}