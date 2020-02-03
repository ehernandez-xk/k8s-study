## Volumes


### emptyDir

Use a init container to clone the a git repository and mount inside a running container

empty Dir is not a persitance volumen only lives on container restarts, if the pod is deleted
or is moved to anoter node the content will be deleted.


```yaml
apiVersion: v1
kind: Pod
metadata:
  name: git-pod
  labels:
    app: git-pod
spec:
  initContainers:
  - name: cloner
    image: alpine
    command: ['sh', '-c', 'apk add --no-cache git && git clone https://github.com/ehernandez-xk/k8s-study.git /myrepo']
    volumeMounts:
    - name: git-repo
      mountPath: /myrepo
  containers:
  - name: reader
    image: eddygt/apphostname:1.0
    volumeMounts:
    - name: git-repo
      mountPath: /myrepo
  volumes:
  - name: git-repo
    emptyDir: {}
```

```bash
 kubectl exec git-pod -- ls /myrepo
```
