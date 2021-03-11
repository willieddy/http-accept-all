package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type emailRequest struct {
	to          *[]string
	cc          *[]string
	bcc         *[]string
	headers     *[]header
	subject     string
	content     string
	flags       string
	contentType string
}

type header struct {
	key   string
	value string
}

type EmailResponse struct {
	Result       string `json:"result"`
	EmailMsgId   string `json:"emailMsgId"`
	ResponseCode int    `json:"responseCode"`
}

func sendEmail() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "sendEmail: ", log.LstdFlags)

		w.Header().Set("Content-Type", "application/json")

		// var request emailRequest
		// err := json.NewDecoder(r.Body).Decode(&request)
		// if err != nil {
		// 	logger.Println("Unable to read request")
		// }
		// logger.Println(request)

		response := EmailResponse{
			Result:       "Success",
			EmailMsgId:   "1",
			ResponseCode: 200,
		}

		logger.Println(response)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.Println(err)
		}
	})
}
