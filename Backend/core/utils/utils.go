package utils

import (
	"net/http"
    "encoding/json"
)

type Responsesss struct {
    Message string `json:"message"`
    Status  string `json:"status"`
    Data    interface{} `json:"data"`
}


func jsonHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "Hello, World!",
        Status:  "Success",

    }

    // Set the Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode the response struct to JSON and write it to the response writer
    json.NewEncoder(w).Encode(response)
}
