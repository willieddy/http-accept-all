package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type emailRequest struct {
	to          []string
	cc          []string
	bcc         []string
	headers     []header
	subject     string
	content     string
	flags       string
	contentType string
}

type header struct {
	key   string
	value string
}

type emailResponse struct {
	result       string
	emailMsgId   string
	responseCode int
}

func sendEmail() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "sendEmail: ", log.LstdFlags)

		w.Header().Set("Content-Type", "application/json")

		var request emailRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			logger.Println("Unable to read request")
		}
		logger.Println(request)

		err = json.NewEncoder(w).Encode(emailResponse{
			result:       "Success",
			emailMsgId:   "1",
			responseCode: 200,
		})
		w.WriteHeader(http.StatusOK)
	})
}
