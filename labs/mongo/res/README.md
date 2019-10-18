# Mongo Cluster

Namespace
kubectl create ns mongo -o yaml --dry-run > namespace.yaml

ConfigMap
kubectl create configmap mongo --from-file=init.sh --dry-run -o yaml > configMap.yaml

Statefulset
Create the StatefulSet I copy paste the basic yaml from the docs
I tested my changes with `kubectl apply -f mongo.yaml --dry-run`
`kubectl explain statefulset.spec.template.spec.volumes.configMap`

Service
kubectl expose statefulset mongo --port 27017 --name mongo --cluster-ip=None --dry-run -o yaml > service.yaml
te previous command does not work because the expose requires the statefulset to be running, but is crashing because the script is waiting for the service
so I create it manually
kubectl create service clusterip mongo --clusterip="None" --tcp=27017 --dry-run -o yaml > service.yaml


Run
kubectl apply -f namespace.yaml
kubectl -n mongo apply -f configMap.yaml
kubectl -n mongo apply -f service.yaml
kubectl -n mongo apply -f mongo.yaml

test

kubectl -n default run --rm -it --generator=run-pod/v1 --image eddygt/apphostname:1.0 mytest sh
host mongo.mongo
mongo.mongo.svc.cluster.local has address 10.244.1.13
mongo.mongo.svc.cluster.local has address 10.244.2.10
mongo.mongo.svc.cluster.local has address 10.244.1.12

host mongo-0.mongo.mongo
mongo-0.mongo.mongo.svc.cluster.local has address 10.244.1.12
host mongo-1.mongo.mongo
mongo-1.mongo.mongo.svc.cluster.local has address 10.244.2.10
host mongo-2.mongo.mongo
mongo-2.mongo.mongo.svc.cluster.local has address 10.244.1.13