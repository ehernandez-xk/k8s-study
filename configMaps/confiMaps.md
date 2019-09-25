## ConfigMaps

are used to provide configuration information for workloads. This can either be fine-grained information (a short string) or a composite value in the form of a file
The key thing is that the ConfigMap is combined with the Pod right before it is run

### From a file

`--from-file`
```bash
# the key is the file name
kubectl create configmap kuard-config --dry-run -o yaml --from-file app.properties > from-file-configmap.yaml
# change the key
kubectl create configmap kuard-config --dry-run -o yaml --from-file=myfile=app.properties > from-file-with-key-configmap.yaml
```

### From a literal
`--from-literal`
```bash
kubectl create configmap kuard-config --dry-run -o yaml --from-file app.properties --from-literal=extra-param=extra-value > from-file-literal-configmap.yaml
# use multiple --from-literal flag if many
```
### From a environment file
`--from-env-file`

## Using a configMap

**Filesystem**
You can mount a ConfigMap into a Pod. A file is created for each entry based on the key name. The contents of that file are set to the value.

**Environment variable**
A ConfigMap can be used to dynamically set the value of an environment variable.

**Command-line argument**
Kubernetes supports dynamically creating the command line for a container based on ConfigMap values.

The following deployment will map the three cases in the pods

```bash
# create the configMap
kubectl apply -f from-file-literal-configmap.yaml
```
Create the basic deploy yaml

```bash
kubectl create deployment my-deploy --image=nginx  --dry-run -o yaml > my-deploy.yaml
```

```bash
# to get some help to config an EnvVar
kubectl explain pod.spec.containers.env.valueFrom.configMapKeyRef
# to get some help to config a file in the file system
kubectl explain pod.spec.containers.volumeMounts
kubectl explain pod.spec.volumes
kubectl explain pod.spec.volumes.configMap
# to add it as command-line argument, just use the EnvVar
kubectl explain pod.spec.containers.command
```


```bash
# check the envVar
kubectl exec <POD-NAME> env | grep MY-CUSTOM-ENV
# check the file
kubectl exec <POD-NAME> cat /usr/config/my-file
# check the command-line argument
kubectl logs <POD-NAME>
```