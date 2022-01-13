# KomposePlayground

## Prerequisite
- Docker
- Minikube
- Helm
- Tilt
- **Put the following into to your .zshrc**

```bash
# Minikube
eval $(minikube docker-env)

# Go
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## Play
1. Run minikube
```bash
minikube start
```
2. Run tilt
```bash
tilt up
```

Try to make change in `main.go`, you will see the change in real time!

## Access url
1. Run tunnel (Doesn't matter if there is no response)
```bash
minikube tunnel --cleanup
```

## Access k8s dashboard
```bash
minikube dashboard
```

## ssh into pod
```bash
kubectl exec --stdin --tty <pod-name> -- /bin/bash
```

## Use redis-cli
```bash
kubectl exec -it <redis-pod> -- redis-cli
```

## Kompose
My helm chart was initially generated by
```bash
kompose convert -c --with-kompose-annotation=false --out kompose-playground
```

## Install PubSub emulator
1. docker build pub-sub-emulator/docker -t gcppubsub:latest
2. helm install pub-sub-emulator pub-sub-emulator

Now it is running!

Below are copy pasted from https://github.com/andrew-jung/gcp-pubsub-emulator

### Create a topic

`PUT http://localhost:8085/v1/projects/test/topics/cats`

`cats` is a new topic name and `test` is the project.

### Get project topics

`GET http://localhost:8085/v1/projects/test/topics`

### Subscribe to topic

`PUT http://localhost:8085/v1/projects/test/subscriptions/test-sub`

`test-sub` is the subscription name

Payload:
```json
{"topic": "projects/test/topics/cats"}
```

### Publish message to topic

`POST http://localhost:8085/v1/projects/test/topics/cats:publish`

Payload:
```json
{"messages": [{"data": "Q2F0IHJ1bGVzIQo="}]}
```
The message is a base 64 encoded string. Original string is `Cat rules!`

### Pull messages from topic

`POST localhost:8085/v1/projects/test/subscriptions/test-sub:pull`

Payload:
```json
{"max_messages": 10}
```
Returned messages are base64 encoded.

## Helm

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

```bash
helm repo add kompose-playground https://quzhi1.github.io/helm-charts
```

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
kompose-playground` to see the charts.

To install the kompose-playground chart:

```bash
helm install kompose-playground kompose-playground/kompose-playground
```

To uninstall the chart:

```bash
  helm delete kompose-playground
```
