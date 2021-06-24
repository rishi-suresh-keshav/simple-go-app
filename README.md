# simple-go-app
A simple-go-app which is a dockerized and deployable to k8

# Make
This project has Makefile support.

```
make build
make pack
make upload
make deploy
make clean
make ship
```

# Local k8 cluster using minikube

Install minikube on macOS
```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
sudo install minikube-darwin-amd64 /usr/local/bin/minikube
```

minikube cluster commands
```
minikube start
minikube pause
minikube stop
```

Start Kubernetes dashboard
```
minkube dashboard
```
