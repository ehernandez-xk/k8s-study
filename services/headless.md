# Headless service

Comparison between **Headless** and **ClusterIP** services

For headless Services, a cluster IP is not allocated, also is useful when proxying is not required and when you don't need load-balancing.
Only consists in set to value None the clusterIP (`service.spec.clusterIP`)

Create a deployment with 3 replicas
```bash
kubectl create deployment myapp --image eddygt/apphostname:1.0
```

```bash
kubectl scale --replicas=3 deployment myapp
```

Create the **Headless** service
```bash
kubectl expose deployment myapp --name myheadless --cluster-ip=None --port=80 --target-port=8080
```

Create the **ClusterIP** service
```bash
kubectl expose deployment myapp --name myclusterip --port=80 --target-port=8080
```

```bash
kubectl get services
NAME          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
kubernetes    ClusterIP   10.96.0.1        <none>        443/TCP   75m
myclusterip   ClusterIP   10.100.178.104   <none>        80/TCP    4s
myheadless    ClusterIP   None             <none>        <none>    52s
```

```bash
kubectl get endpoints
NAME          ENDPOINTS                                         AGE
kubernetes    172.17.0.3:6443                                   75m
myclusterip   10.244.1.2:8080,10.244.1.3:8080,10.244.2.2:8080   11s
myheadless    10.244.1.2,10.244.1.3,10.244.2.2                  59s
```

```bash
kubectl proxy
Starting to serve on 127.0.0.1:8001
```

Review if the services are responding
```bash
http://localhost:8001/api/v1/namespaces/default/services/myclusterip/proxy/
http://localhost:8001/api/v1/namespaces/default/services/myheadless/proxy/
```

I create a pod to run some DNS discovery. The image has already installed curl and bind-tools
```bash
kubectl run --generator=run-pod/v1 --rm -it --image eddygt/apphostname:1.0 mytest sh
```

```bash
# headless services returns the pod's ips
host myheadless
myheadless.default.svc.cluster.local has address 10.244.2.2
myheadless.default.svc.cluster.local has address 10.244.1.2
myheadless.default.svc.cluster.local has address 10.244.1.3
# clusterip service returns the service's ip
host myclusterip
myclusterip.default.svc.cluster.local has address 10.100.178.104
```

```bash
# access to the pod port
for i in $(seq 1 3) ; do curl myheadless:8080; done
```
```bash
# access to the service port
for i in $(seq 1 3) ; do curl myclusterip; done
```

## Headless service for Statefulset

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: nginx # has to match .spec.template.metadata.labels
  serviceName: "nginx"
  replicas: 3 # by default is 1
  template:
    metadata:
      labels:
        app: nginx # has to match .spec.selector.matchLabels
    spec:
      terminationGracePeriodSeconds: 10 # Optional duration in seconds the pod needs to terminate gracefully
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80
          name: web
```

```bash
kubectl get pods
NAME    READY   STATUS    RESTARTS   AGE
web-0   1/1     Running   0          2m12s
web-1   1/1     Running   0          2m2s
web-2   1/1     Running   0          114s
```

The domain managed by this Service takes the form: $(service name).$(namespace).svc.cluster.local, where “cluster.local” is the cluster domain.
As each Pod is created, it gets a matching DNS subdomain, taking the form: $(podname).$(governing service domain), where the governing service is defined by the serviceName field on the StatefulSet.

```bash
kubectl run --rm -it --generator=run-pod/v1 --image eddygt/apphostname:1.0 mytest sh
```

```bash
host nginx
nginx.default.svc.cluster.local has address 10.244.1.4
nginx.default.svc.cluster.local has address 10.244.2.4
nginx.default.svc.cluster.local has address 10.244.1.3
```

```bash
host web-0.nginx
web-0.nginx.default.svc.cluster.local has address 10.244.1.3
```

```bash
# Test the domain
curl web-0.nginx
curl web-0.nginx.default.svc.cluster.local
```

### Pod Name Label
When the StatefulSet controller creates a Pod, it adds a label, statefulset.kubernetes.io/pod-name, that is set to the name of the Pod. This label allows you to attach a Service to a specific Pod in the StatefulSet.

```bash
kubectl describe pod web-0
```

```
Labels:     app=nginx
            controller-revision-hash=web-7858455875
            statefulset.kubernetes.io/pod-name=web-0
```
