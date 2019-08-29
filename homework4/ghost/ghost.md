## Ghost

https://ghost.org/

Ghost is a popular blogging engine with a clean interface written in JavaScript. It can either use a file-based SQLite database or MySQL for storage.

### Create the ConfigMap
```bash
kubectl create cm --from-file ghost-config.js ghost-config
```

### The deployment
```bash
kubectl apply -f ghost.yaml
```

### Create the Service
```bash
kubectl expose deployments ghost --port=2368
```
### Proxy to localhost
```bash
kubectl proxy
```

Then visit http://localhost:8001/api/v1/namespaces/default/services/ghost/proxy/ in your web browser to begin interacting with Ghost.

