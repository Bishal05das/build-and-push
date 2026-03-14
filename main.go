package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type:","application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode("hello world!")
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /",Greet)
	fmt.Println("Server listening on port 3000")
	err := http.ListenAndServe(":3000",mux)
	if err != nil {
		fmt.Println("error in server lisening: ",err)
	}
}