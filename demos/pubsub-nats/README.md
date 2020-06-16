# Dapr Pub/Sub messaging with NATS (nats.io)

Replace NATS info in `components/nats.yaml` (this example uses the public NATS server at `demo.nats.io` which you can continue with for the purposes of this demo)

```bash`
export APP_PORT=8080
export NATS_SUBJECT=demo
```

Run the app

```bash
dapr run --app-id natsapp --port 3500 --app-port $APP_PORT --components-path components go run nats.go
```

Check NATS connections. For the public NATS instance, you can browse to `http://demo.nats.io:8222/connz` and confirm that your client app is indeed connected to NATS - scroll down to the bottom of the page, you client app will likely be the last entry and you can confirm it by seeing the `uptime` which should be in the order of a few seconds and `in_msgs`, `out_msgs` which should be `0` at this point

Use Dapr CLI to publish messages.. You can also try sending messages directly using the Dapr Pub/Sub endpoint (using any HTTP client such as `curl`)

```bash
dapr publish -t demo -d foo
curl -X POST http://localhost:3500/v1.0/publish/demo -d 'foooo'
```