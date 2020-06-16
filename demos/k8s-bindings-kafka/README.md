# Dapr Kafka (using Event Hubs) Bindings on Kubernetes

Deploy Bindings component - replace Event Hubs info in `deploy/kafka.yaml` first

```bash
kubectl apply -f deploy/kafka.yaml
kubectl get component.dapr.io
```


Deploy Dapr consumer application

```bash
kubectl apply -f deploy/app.yaml
kubectl get pods -l=app=kafkaapp
```

Check logs

```bash
kubectl logs <POD> -c daprd
kubectl logs -f <POD> -c kafkaapp
```

Start producer app after replacing Event Hubs values in `start-producer.sh` file

Scale out the Dapr app

```bash
kubectl scale deployment/kafkaapp --replicas=2
```

To remove the Binding component and app:

```bash
kubectl delete -f deploy
```