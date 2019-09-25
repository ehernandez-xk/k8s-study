# Kind

https://kind.sigs.k8s.io

Is a tool for running local Kubernetes clusters using Docker container as “nodes”.

## Usage

The `create` by default creates 1 node but you can pass a configuration to specify the number of pods
you can create more than one cluster. The default cluster name is `kind`
```bash
kind create cluster --config kind-cluster-3-nodes.yaml --name=mycluster
```

There are more options you can configure here
```yaml
# kind-cluster-3-nodes.yaml
# three node (two workers) cluster config
kind: Cluster
apiVersion: kind.sigs.k8s.io/v1alpha3
nodes:
- role: control-plane
- role: worker
- role: worker
```

Configure the kubectl command to the cluster
```bash
export KUBECONFIG="$(kind get kubeconfig-path --name=mycluster)"
```

```bash
kubectl get nodes
NAME                      STATUS   ROLES    AGE   VERSION
mycluster-control-plane   Ready    master   23m   v1.15.3
mycluster-worker          Ready    <none>   22m   v1.15.3
mycluster-worker2         Ready    <none>   22m   v1.15.3
```

Currently kind does not support private registries but we can load the image to all the nodes
```bash
kind load docker-image nginx --name=mycluster
```

```bash
kind delete cluster --name=mycluster
```
