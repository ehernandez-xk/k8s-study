# Minikube

Install
```bash
brew cask install minikube
```

### Install the Hyperkit driver
```bash
brew install docker-machine-driver-hyperkit
sudo chown root:wheel /usr/local/opt/docker-machine-driver-hyperkit/bin/docker-machine-driver-hyperkit
sudo chmod u+s /usr/local/opt/docker-machine-driver-hyperkit/bin/docker-machine-driver-hyperkit
```

### Start the cluster
```bash
minikube start --vm-driver hyperkit --memory 6144 --cpus 4
```

## Output the deployment
```bash
kubectl create deployment nginx --image nginx --dry-run -o yaml > nginx/deployment.yaml
```

## Output the service
```bash
kubectl expose -f nginx/deployment.yaml --port=80 --type=LoadBalancer --dry-run -o yaml > nginx/service.yaml
# or
kubectl expose -f nginx/deployment.yaml --port=80 --type=NodePort --dry-run -o yaml > nginx/service.yaml
```

### Review and clean the yaml files

## Deploy to cluster
```bash
kubectl apply -f nginx/deployment.yaml
kubectl apply -f nginx/service.yaml
```

## Expose the service
```bash
minikube service nginx
```

### View on localhost
```bash
minikube dashboard --url
http://127.0.0.1:58606/api/v1/namespaces/kube-system/services/http:kubernetes-dashboard:/proxy/
```
