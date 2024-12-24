package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OperationRequest struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type OperationResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var req OperationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := req.Number1 + req.Number2
	resp := OperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	var req OperationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := req.Number1 - req.Number2
	resp := OperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	var req OperationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := req.Number1 * req.Number2
	resp := OperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	var req OperationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if req.Number2 == 0 {
		http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
		return
	}

	result := req.Number1 / req.Number2
	resp := OperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/add", loggingMiddleware(addHandler))
	http.HandleFunc("/subtract", loggingMiddleware(subtractHandler))
	http.HandleFunc("/multiply", loggingMiddleware(multiplyHandler))
	http.HandleFunc("/divide", loggingMiddleware(divideHandler))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
