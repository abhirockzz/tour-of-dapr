package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string
var bindingName string

func init() {
	port = os.Getenv("APP_PORT")
	if port == "" {
		log.Fatalf("missing environment variable %s", "APP_PORT")
	}

	bindingName = os.Getenv("BINDING_NAME")
	if port == "" {
		log.Fatalf("missing environment variable %s", "BINDING_NAME")
	}
}

func main() {
	http.HandleFunc("/"+bindingName, func(rw http.ResponseWriter, req *http.Request) {
		var _time TheTime
		err := json.NewDecoder(req.Body).Decode(&_time)
		if err != nil {
			fmt.Println("error reading message from kafka binding", err)
			rw.WriteHeader(500)
			return
		}
		fmt.Printf("time from Event Hubs '%s'\n", _time.Time)
	})
	http.ListenAndServe(":"+port, nil)
}

type TheTime struct {
	Time string `json:"time"`
}
