package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response http response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receive request from %s", r.RemoteAddr)
		responseBody, err := json.MarshalIndent(&Response{
			Code:    0,
			Message: "ok",
		}, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(responseBody)
	})
	http.ListenAndServe(":8080", nil)
}
