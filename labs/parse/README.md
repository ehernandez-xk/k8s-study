## Parse

https://parseplatform.org/

The Parse server is a cloud API dedicated to providing easy-to-use storage for mobile applications. It provides a variety of different client libraries that make it easy to integrate with Android, iOS, and other mobile platforms. Parse was purchased by Facebook in 2013 and subsequently shut down. Fortunately for us, a compatible server was open sourced by the core Parse team and is available for us to use. This section describes how to set up Parse in Kubernetes.

### Prerequisites

- Mongo

This section assumes you have a three-replica Mongo cluster running in Kubernetes with the names mongo-0.mongo, mongo-1.mongo, and mongo-2.mongo.

These instructions also assume that you have a Docker login; if you don’t have one, you can get one for free at https://docker.com

**Building the parse-server**
The open source parse-server comes with a Dockerfile by default, for easy contain‐ erization. First, clone the Parse repository:
```bash
git clone https://github.com/ParsePlatform/parse-server
# Then move into that directory and build the image:
cd parse-server
docker build -t ${DOCKER_USER}/parse-server .
# Finally, push that image up to the Docker hub:
docker push ${DOCKER_USER}/parse-server
```

### Steps

1. deploy the parse-server with the following configuration
- EnvVars:

```yaml
env:
- name: PARSE_SERVER_DATABASE_URI
    value: "mongodb://mongo-0.mongo:27017,\
    mongo-1.mongo:27017,mongo-2.mongo\
    :27017/dev?replicaSet=rs0"
- name: PARSE_SERVER_APP_ID
    value: my-app-id
- name: PARSE_SERVER_MASTER_KEY
    value: my-master-key
```

2. Create the service for the deployment 
3. use the label: `run: parse-server`
4. expose the deployment using a Service (the container exposes port `1337`)
