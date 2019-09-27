
### Create the namespace
```bash
kubectl create namespace jupyter --dry-run -o yaml > namespace.yaml
kubectl apply -f namespace.yaml
```

### Change namespace (optional)
Help us to change the current namespace
```bash
kubectl config set-context --current --namespace=jupyter
```

### Create the deployment
```bash
kubectl -n jupyter create deploy --image jupyter/scipy-notebook:abdb27a6dfbb jupyter --dry-run -o yaml > deployment.yaml
kubectl apply -f deployment.yaml
```
I added the Always restartPolicy and the desired label
`kubectl explain pod.spec.restartPolicy`

### Create the service
```bash
kubectl -n jupyter expose deployment/jupyter --name jupyter --port 80 --target-port=8888 --dry-run -o yaml > service.yaml
kubectl apply -f service.yaml
```
I add NodePort type
`kubectl explain service.spec.type`

### Checking
```bash
kubectl -n jupyter get endpoints
kubectl -n jupyter logs jupyter-5cc4fbd5b-wdhg4
# from logs ?token=f68b263257289cd9416c71603f3ed185d5d7d7979c846b8e
# to see in localhost:8888
kubectl -n jupyter port-forward jupyter-5cc4fbd5b-wdhg4 8888
```

### Expose
To see if the services was correctly configured
```bash
minikube service list
minikube service jupyter
```

### Ping the service
from default namespace
```bash
kubectl config set-context --current --namespace=default
kubectl run --rm -it --generator=run-pod/v1 --image eddygt/apphostname:1.0 mytest sh
host jupyter.jupyter
jupyter.jupyter.svc.cluster.local
curl -v jupyter.jupyter.svc.cluster.local:8888/?token=f68b263257289cd9416c71603f3ed185d5d7d7979c846b8e
```
