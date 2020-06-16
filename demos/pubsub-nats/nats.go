package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var port string
var natsSubject string

func init() {
	port = os.Getenv("APP_PORT")
	if port == "" {
		log.Fatalf("missing env var %s", "APP_PORT")
	}

	natsSubject = os.Getenv("NATS_SUBJECT")
	if natsSubject == "" {
		log.Fatalf("missing env var %s", "NATS_SUBJECT")
	}
}

func main() {
	http.HandleFunc("/dapr/subscribe", func(w http.ResponseWriter, r *http.Request) {
		sub := Subscription{Topic: natsSubject, Route: natsSubject}
		subs := []Subscription{sub}
		json.NewEncoder(w).Encode(&subs)
		fmt.Println("subscribed to NATS subject", natsSubject)
	})

	http.HandleFunc("/"+natsSubject, func(w http.ResponseWriter, r *http.Request) {
		payload, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		fmt.Println("Payload from NATS - " + string(payload))
		var event Event

		json.NewDecoder(bytes.NewReader(payload)).Decode(&event)
		fmt.Println("Event ID - " + event.ID)
		fmt.Println("Event Source - " + event.Source)
		fmt.Println("Event Data - " + event.Data)
	})

	fmt.Println("starting HTTP server....")
	http.ListenAndServe(":"+port, nil)
}

type Subscription struct {
	Topic string `json:"topic"`
	Route string `json:"route"`
}

type Event struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Data   string `json:"data"`
}
