## Parse

The Parse server is a cloud API dedicated to providing easy-to-use storage for mobile applications. It provides a variety of different client libraries that make it easy to integrate with Android, iOS, and other mobile platforms. Parse was purchased by Facebook in 2013 and subsequently shut down. Fortunately for us, a compatible server was open sourced by the core Parse team and is available for us to use. This section describes how to set up Parse in Kubernetes.

### Prerequisites

- Mongo

### Install

Login to Docker to use my public images in the registry

```bash
docker login
# user eddygt
```

```bash
git clone https://github.com/ParsePlatform/parse-server
cd parse-server
docker build -t ${DOCKER_USER}/parse-server .
docker push ${DOCKER_USER}/parse-server
```

### Proxy to localhost
```bash
kubectl proxy
```

Then visit http://localhost:8001/api/v1/namespaces/default/services/parse-server/proxy/ in your web browser to begin interacting with Ghost.
