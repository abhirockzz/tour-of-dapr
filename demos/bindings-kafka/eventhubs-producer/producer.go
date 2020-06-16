package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go/v2"
)

var eventHubNamespace string
var eventHubKey string
var eventHubName string

const eventHubConnectionStringFormat = "Endpoint=sb://%s.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=%s;EntityPath=%s"

func init() {
	eventHubNamespace = os.Getenv("EVENT_HUBS_NAMESPACE")
	eventHubKey = os.Getenv("EVENT_HUBS_KEY")
	eventHubName = os.Getenv("EVENT_HUB_NAME")

	if eventHubNamespace == "" || eventHubKey == "" || eventHubName == "" {
		log.Fatalf("missing required environment variables")
	}
}

func main() {
	eventHubConnectionString := fmt.Sprintf(eventHubConnectionStringFormat, eventHubNamespace, eventHubKey, eventHubName)

	hub, err := eventhub.NewHubFromConnectionString(eventHubConnectionString)
	if err != nil {
		log.Fatalf("failed to create hub client\n\n%s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer hub.Close(ctx)
	defer cancel()
	if err != nil {
		log.Fatalf("failed to get hub %s\n", err)
	}

	format := "Mon Jan _2 15:04:05 2006"
	ctx = context.Background()

	for {
		_time := time.Now().Format(format)
		text := `{"time":"` + _time + `"}`
		err := hub.Send(ctx, eventhub.NewEventFromString(text))
		if err != nil {
			fmt.Printf("Error sending msg: %s\n", err)
		} else {
			fmt.Printf("Sent message %s\n", text)
		}
		time.Sleep(2 * time.Second)
	}
}
