```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: tenant-max-mem
  namespace: default
spec:
  limits:
  - max:
      memory: 250Mi
    type: Container
```

following will fail because it excedess the limit of the namespace
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: too-much-memory
  namespace: default 
spec:
  containers:
  - name: too-much-mem
    image: nginx
    resources:
      requests:
        memory: "300Mi"
```

