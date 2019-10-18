## Mongo

Create a StatefulSet-based MongoDB cluster.

Steps:
1. Create a new namespace `mongo`
2. Create a Statefulset for MongoDB to have the pods mongo-0, mongo-1 and mongo-2
    - 3 replicas
    - match label `app: mongo`
    - `image: mongo:3.4.1`
    - command `mongod --replSet rs0`
    - containerPort 27017
    - mount in volumes the init.sh files as configMap (name: mongo-init)
3. A init container is necessary to initialize the server (this will run a init.sh script)
    - `image mongo:3.4.1`
    - command `bash /config/init.sh`
    - we need a volumen to have the init script (name: config, mountPath: /config)
4. Create a headless service for mongo cluster (is a service without Cluster IP address `clusterIP: None`)
    - name mongo
    - port 27017
    - selector `app: mongo`
    - after the creation review the DNS entries that were created for this services (should be 4)
5. Create a configMap to holds de following initializing script and add it using the init.sh key
    - name `mongo-init`

Now that you have all in place you can create the cluster with:
```bash
kubectl apply -f ns.yaml
kubectl apply -f configMap.yaml
kubectl apply -f service.yaml
kubectl apply -f mongo.yaml
```

init.sh file
```bash
#!/bin/bash
# Need to wait for the readiness health check to pass so that the
# mongo names resolve. This is kind of wonky.
until ping -c 1 ${HOSTNAME}.mongo; do
    echo "waiting for DNS (${HOSTNAME}.mongo)..."
sleep 2 done
until /usr/bin/mongo --eval 'printjson(db.serverStatus())'; do
    echo "connecting to local mongo..."
    sleep 2
done
echo "connected to local."
HOST=mongo-0.mongo:27017
until /usr/bin/mongo --host=${HOST} --eval 'printjson(db.serverStatus())'; do
    echo "connecting to remote mongo..."
    sleep 2
done
echo "connected to remote."
if [[ "${HOSTNAME}" != 'mongo-0' ]]; then
    until /usr/bin/mongo --host=${HOST} --eval="printjson(rs.status())" \
        | grep -v "no replset config has been received"; do
    echo "waiting for replication set initialization"
    sleep 2
    done
    echo "adding self to mongo-0"
    /usr/bin/mongo --host=${HOST} \
        --eval="printjson(rs.add('${HOSTNAME}.mongo'))"
fi
if [[ "${HOSTNAME}" == 'mongo-0' ]]; then
    echo "initializing replica set"
    /usr/bin/mongo --eval="printjson(rs.initiate(\
        {'_id': 'rs0', 'members': [{'_id': 0, \
        'host': 'mongo-0.mongo:27017'}]}))"
fi
echo "initialized"
while true; do
    sleep 3600
done
```

```bash
kubectl apply -f mongo-config-map.yaml
kubectl apply -f mongo-service.yaml
kubectl apply -f mongo.yaml
# ping
kubectl run -it --rm --image busybox busybox ping mongo-1.mongo
```

The DNS resolution for this statefulSet mongo is `mongo-0⁠.mongo⁠.default⁠.svc⁠.cluster​.local`
