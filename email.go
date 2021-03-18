package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EmailRequest struct {
	To          *[]string `json:"to"`
	Cc          *[]string `json:"cc"`
	Bcc         *[]string `json:"bcc"`
	Headers     *[]Header `json:"headers"`
	Subject     string    `json:"subject"`
	Content     string    `json:"content"`
	Flags       string    `json:"flags"`
	ContentType string    `json:"contentType"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EmailResponse struct {
	Result       string `json:"result"`
	EmailMsgId   string `json:"emailMsgId"`
	ResponseCode int    `json:"responseCode"`
	DeliveryTime string `json:"deliveryTime"`
}

func sendEmail() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		time.Sleep(100 * time.Millisecond)

		var request EmailRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println("Unable to read request")
		}
		fmt.Println(request)

		response := EmailResponse{
			Result:       "Success",
			EmailMsgId:   "1",
			ResponseCode: 200,
			DeliveryTime: "0000",
		}

		fmt.Println(response)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Println(err)
		}
	})
}
