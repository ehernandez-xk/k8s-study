## Secrets

The easiest way to create a secret or a ConfigMap is via kubectl create secret generic or kubectl create configmap. There are a variety of ways to specify the data items that go into the secret or ConfigMap. These can be combined in a single command:

--from-file=<filename>
Load from the file with the secret data key the same as the filename.

--from-file=<key>=<filename>
Load from the file with the secret data key explicitly specified.

--from-file=<directory>
Load all the files in the specified directory where the filename is an acceptable key name.

--from-literal=<key>=<value>
Use the specified key/value pair directly.

### Create and Update

```bash
kubectl create secret generic kuard-tls \
  --from-file=kuard.crt --from-file=kuard.key \
  --dry-run -o yaml | kubectl replace -f -
```

```bash
kubectl edit configmap my-config
```
Note:
you could also do this with a secret, but you’d be stuck managing the base64 encoding of values on your own

Once a ConfigMap or secret is updated using the API, it’ll be automatically pushed to all volumes that use that ConfigMap or secret
