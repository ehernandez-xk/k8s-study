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

# Digital Ocean Config

Create the cluster using UI interface
## configure kubectl to use the kubeconfig file

### Using --kubeconfig flag
```bash
kubectl --kubeconfig=/Users/eddhernandez/Desktop/k8s-1-14-3-do-0-sfo2-1562260891246-kubeconfig.yaml --context do-sfo2-k8s-1-14-3-do-0-sfo2-1562260891246 get nodes
```

### Change default config file
Add a default for the kubectl
```bash
export KUBECONFIG=~/Desktop/k8s-1-14-3-do-0-sfo2-1562260891246-kubeconfig.yaml
```

### Using many config files
```bash
export KUBECONFIG=~/.kube/config:~/Desktop/k8s-1-14-3-do-0-sfo2-1562260891246-kubeconfig.yaml
```

Now you can see the contexts

```bash
kubectl config get-contexts

# now minikube is the current
CURRENT   NAME                                         CLUSTER                                      AUTHINFO                                           NAMESPACE
          do-sfo2-k8s-1-14-3-do-0-sfo2-1562260891246   do-sfo2-k8s-1-14-3-do-0-sfo2-1562260891246   do-sfo2-k8s-1-14-3-do-0-sfo2-1562260891246-admin
          docker-for-desktop                           docker-for-desktop-cluster                   docker-for-desktop
*         minikube                                     minikube                                     minikube

# change the contexts
kubectl config use-context do-sfo2-k8s-1-14-3-do-0-sfo2-1562260891246
```

### Deploy to cluster
```bash
kubectl apply -f nginx/deployment.yaml
kubectl apply -f nginx/service.yaml
```

### Get Load Balancer
A new EXTERNAL-IP for the Load Balancer is created
```bash
kubectl get services
```
