## Guestbook application

This is an stateless application whit redis and PHP

Tutorial from https://kubernetes.io/docs/tutorials/stateless-application/guestbook/
```bash
kubectl apply -f redis-master-deploy.yaml
kubectl apply -f redis-master-service.yaml
kubectl apply -f redis-slave-deployment.yaml
kubectl apply -f redis-slave-service.yaml
kubectl apply -f frontend-deployment.yaml
kubectl apply -f frontend-service.yaml
```

