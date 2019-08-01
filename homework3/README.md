
## Replicaset

```bash
kubectl get replicaset kuard-health-checks-6b969559fc -o yaml > kuard-replicaset.yaml
```

Clean up the yaml and apply the replica set

```bash
kubectl apply -f kuard-replicaset.yaml
```

See that the replicaset does not add new labels to the pod

```bash
kubectl get pods --show-labels
```

Replicaset add to each pod an Object to identifiy that this belong to a replicaset `pod.metadata.ownerReferences`

### Replicate

Imperative

```bash
kubectl scale replicasets kuard --replicas=4
```
Declarative
```bash
# Edit the replicaset yaml file and apply the changes
```

## horizontal pod autoscaling (HPA)

It is a control loop that automatically scales the number of pods in a replication controller (deployment or replicaset). By default check period is 15 seconds. It can be changed with `--horizontal-pod-autoscaler-sync-period` flag on the controll manager startup

Queries the resource utilization against the metrics specified in each HPA definition

HPA requires [metrics-server](https://github.com/kubernetes-incubator/metrics-server)

Minikube v1.1.1 comes this addon disabled
```bash
minikube addons list | grep metrics
```

```bash
git clone https://github.com/kubernetes-incubator/metrics-server.git
kubectl create -f deploy/1.8+/
```

For minikube I need to do the following to disable tls

```bash
kubectl edit deploy -n kube-system metrics-server
## add two flags and the args array in the metrics-server container
    args:
    - --kubelet-insecure-tls
    - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
```

### Autoscale

```bash
kubectl autoscale rs kuard --min=2 --max=5 --cpu-percent=80
kubectl get horizontalpodautoscalers
kubectl get hpa/kuard
```

Delete 
```bash
# Delete only the replicaset but not the pods
kubectl delete rs kuard --cascade=false
# Missing pods
kubectl delete pod/kuard-vz7tk pods/kuard-wtqcz
kubectl delete pods kuard-x8wbz kuard-xmmwl
```

## DaemonSets

```bash
# get an initial yaml
kubectl -n kube-system get daemonset/kube-proxy -o yaml > daemonset/kube-daemonset.yaml
```
Review and cleanup the yaml then apply the daemonset
```bash
kubectl apply -f daemonset/kube-daemonset.yaml
```

### Limiting DaemonSets to Specific Nodes

Deamonset by default creates a pod on each node in the cluster, but this can be chanded by limiting a subset of nodes where you want your pods

Add some labels to the node
```bash
kubectl label nodes minikube ssd=true
```
Review the labels of the node
```bash
kubectl get nodes --show-labels
kubectl get node --selector ssd=true
```

Add the follwoing nodeSelector map in the spec

`daemonset.spec.template.spec.nodeSelector`
```bash
spec:
    nodeSelector:
      ssd: "true"
    containers:
    - image: gcr.io/kuar-demo/kuard-amd64:blue
```
```bash
kubectl apply -f daemonset/kube-daemonset-selector.yaml 
```

Remove the node's label to see if the daemonset remove the pod from that node

```bash
kubectl label nodes minikube --overwrite ssd=false
```

### Update a deamonset
You can delete the daemonset with the flag `--cascade=false` this won't destroy de pods only the parent object. Then create the new daemontset and this will reconize the old pods without touch them. To start the new pods with the new template you will need to delete one by one the pods and the deamonset will start creating the new pods.

```bash
PODS=$(kubectl get pods -o jsonpath --template='{.items[*].metadata.name}')
for x in $PODS; do
  kubectl delete pods ${x}
  sleep 60
done
```

