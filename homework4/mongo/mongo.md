## Mongo

```bash
kubectl apply -f mongo-config-map.yaml
kubectl apply -f mongo-service.yaml
kubectl apply -f mongo.yaml
# ping
kubectl run -it --rm --image busybox busybox ping mongo-1.mongo
```

The DNS resolution for this statefulSet mongo is `mongo-0⁠.mongo⁠.default⁠.svc⁠.cluster​.local`
