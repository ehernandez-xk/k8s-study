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
