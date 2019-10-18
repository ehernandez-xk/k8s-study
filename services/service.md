# Services

When a typical service is created, an IP address is also created and the DNS service is populated with an A Record that points to that IP adress
When you create a service of ExternalName, the DNS service istead pupulate a CNAME record that points to the external name you especified (e.g database.company.com)

### External Name

This only works for DNS address and not when you have an IP, (see below for an external service with an IP)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: google
spec:
  externalName: google.com
  type: ExternalName
```

```bash
kubectl run --rm -it --generator=run-pod/v1 --image eddygt/apphostname:1.0 mytest sh
curl google.default
# or
curl google.default.svc.default.cluster
```

### External service using an IP

Create a service wihtout a label selector, but also without the ExternalName type used above

```yaml
kind: Service
apiVersion: v1
metadata:
  name: external-ip-database
```
with this service a virtual IP address for this service is created also an A record is populated
Becuase there is no labels seclectos no enpoits are created, so we need to create our endpoints

```yaml
kind: Endpoints
apiVersion: v1
metadata:
  name: external-ip-database
subsets:
  - addresses:
    - ip: 192.168.0.1
    ports:
    - port: 3306
```
