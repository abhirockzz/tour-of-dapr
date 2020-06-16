# Kafka Input Bindings using Azure Event Hubs

Replace Event Hubs Kafka broker info in `components/kafka.yaml`

```bash
export APP_PORT=8080
export BINDING_NAME=kafka-input
```

Run the app

```bash
dapr run --app-id kafkaapp --app-port $APP_PORT --components-path components go run app.go 
```

To run producer app, replace Event Hubs info in `start-producer.sh`

```bash
./start-producer.sh
```