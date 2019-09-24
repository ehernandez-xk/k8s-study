## Role and ClusterRole

In the RBAC API, a role contains rules that represent a set of permissions. Permissions are purely additive (there are no “deny” rules). A role can be defined within a namespace with a Role, or cluster-wide with a ClusterRole

Role for namespace permissions
ClusterRole for any particular namespace (cluster-scoped)


To enable RBAC, start the apiserver with --authorization-mode=RBAC


```bash
# create a Role named "pod-reader" that allows user to perform "get", "watch" and "list" on pods
kubectl create role pod-reader --verb=get --verb=list --verb=watch --resource=pods --dry-run -o yaml > role-test.yaml
# or
kubectl create role pod-reader --verb=get,list,watch --resource=pods --dry-run -o yaml > role-test.yaml
```

Role wihtout any permission

```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: lockdown
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: [""]
  verbs: [""]
```

## Service Account

```bash
kubectl create serviceaccount sa-test --dry-run -o yaml > service-account.yaml
```

## Rolebinding

```bash
kubectl create rolebinding rb-test --role=pod-reader --serviceaccount=default:sa-test --dry-run -o yaml > rb-test.yaml
```

## Checking permissions

```bash
kubectl auth can-i create pods --all-namespaces
yes
```
```bash
kubectl auth can-i get pods --namespace default --as system:serviceaccount:default:sa-test
yes
```

## Deployment with permissions

```bash
kubectl create deployment my-dep --image=nginx:1.13 --dry-run -o yaml > nginx-deploy.yaml
```
`kubectl explain deployment.spec.template.spec.serviceAccountName`

## Nginx without permissions

```bash
kubectl apply -f nginx-lockdown.yaml
kubectl auth can-i get pods --namespace lockdown --as system:serviceaccount:lockdown:sa-lockdown
no
```