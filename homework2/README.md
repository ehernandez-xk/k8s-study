## Health Checks

Output a sample yaml

```bash
kubectl create deployment kuard-health-checks --image gcr.io/kuar-demo/kuard-amd64:blue --dry-run -o yaml > health-checks/kuard-livenessProbe.yaml
```

Clean the yaml and I added the port and livenessProbe seccion


### LivenessProbe

```bash
kubectl apply -f health-checks/kuard-livenessProbe.yaml
```
When fails the container is restarted

To test it using the UI

```bash
kubectl port-forward <a-pod-name> 8080
```

### RedinessProbe

```bash
kubectl apply -f kuard-redinessProbe.yaml
```
If pod is READY is assigned to the endpoints in the service. And it is consider READY if the redinessProbe passed

```bash
kubectl port-forward <a-pod-name> 8080
```

```bash
# See endpoints
kubectl get endpoints
```

## Resources Request

Output the deployment

```bash
kubectl create deployment kuard-resources --image gcr.io/kuar-demo/kuard-amd64:blue --dry-run -o yaml > resources-request/kuard-resources.yaml
```

The resources config goes in `deployment.spec.template.spec.containers.resources`

### Minimum Required Resources

The minimal resources required to schedule the pods. The resources are configured per pods, this is because you can have sidecars that may required different config.

The scheduler does a sum of the container resources to know where to schedule the pod in a node.

### Request Limit Details

For this limit the kubelet kill the container if the limit is reached

```bash
kubectapply -f resources-request/kuard-resources.yaml
```
```bash
# start whatching the pods to see the kill
kubectl get pods -w
``` 

Open a new terminal to see the `OOMKilled` then forward the pod to increase the memory allocation. The deployment will create a new container.


## Volumes

`deployment.spec.template.spec.volumes`

Volumes definitions that may be mounted by the containers, the name of the volumen and the host path where this whill be mounted. If a pod is deleted or container restarts the data in the filesystem is deleted.

`pod.spec.containers.volumeMounts`

The definition inside the containers.volumeMounts is where you mount a volume defined in the spec, using the name and the path to mount inside the container

```bash
kubectl apply -f volume-mounts/kuard-volumes.yaml
```

To test it

```bash
minikube ssh
touch /var/lib/kuard/filetest
exit # exit the minikube terminal
kubectl port-forward kuard-volumes-865d648846-z5wzn 8080 # see the filesystem under /data
```

## Labels

Some useful commands

```bash
kubectl run alpaca-prod \
  --image=gcr.io/kuar-demo/kuard-amd64:blue \
  --replicas=2 \
  --labels="ver=1,app=alpaca,env=prod"
```
```bash
# show objects labels
kubectl get deployments --show-labels
kubectl get pods --show-labels
# add/modify label of the deployment (not labels of the template)
kubectl label deployments alpaca-test "canary=true"
# see deployments and adds shows specific label
kubectl get deployments -L canary
# remove a label of a deployment
kubectl label deployments alpaca-test "canary-"
```

```bash
# matching with label selectors
kubectl get pods --selector="ver=2"
kubectl get pods --selector="app=bandicoot,ver=2"
kubectl get pods --selector="app in (alpaca,bandicoot)"
kubectl get deployments --selector="canary"
```
```bash
# Other match options
key=value       key is set to value
key!=value      key is not set to value
key             key is set
!key            key is not set

key in (value1, value2)     key is one of value1 or value2
key notin (value1, value2)  key is not one of value1 or value2
```

## Services

DNS resolution A record for a service

```bash
# FQDN
<service-name>.<namespace>.svc.cluster.local
alpaca-prod.default.svc.cluster.local.
```
**alpaca-prod** The name of the service in question.
**default** The namespace that this service is in.
**svc** Recognizing that this is a service. This allows Kubernetes to expose other types of things as DNS in the future.
**cluster.local.** The base domain name for the cluster. This is the default and what you will see for most clusters. Administrators may change this to allow unique DNS names across multiple clusters.

### Calling a service
<service-name> in the same namespace where calling it
<service-name>.<namespace> in different namespace
FQDN where you want

### See endpoints of a deployment

```bash
kubectl get endpoints alpaca-prod --watch
```

### Access to the service

Tunneling ssh to a NodePort:
If the service is ClusterIP you can edit using `kubectl edit service <service-name>`. See port assigned to the service

```bash
ssh <node> -L 8080:localhost:32711
```
For minikube you can access directly using the IP that minikube is using `minikube status` and then add the port assigned

## Manual Service Discovery

```bash
kubectl get pods -o wide --show-labels
kubectl get pods -o wide --selector=app=alpaca,env=prod
```
