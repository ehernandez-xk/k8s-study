# Jupyter Lab

**Credits:** From book Kubernentes Up And Running

https://jupyter.org/

The Jupyter Project is a web-based interactive scientific notebook for exploration and visualization. It is used by students and scientists around the world to build and explore data and data visualizations. Because it is both simple to deploy and interest‐ ing to use, it’s a great first service to deploy on Kubernetes.

### Objetive:
Deploy the application with 1 replica to see in the browser the Jupyter dasbord, Create the app in a new namespace and to expose the deployment using a Service
Save all the objects you create in yaml files the idea is to deploy the full app with a single apply eg  `kubectl apply -f jupyter-k8s/`

### Steps:
1. Create a new namespace
2. Deploy the app using a Deployment
    - 1 replica
    - Use the image `jupyter/scipy-notebook:abdb27a6dfbb` (the images is quite large, use watch to wait for the Pods)
    - Use the label `run: jupyter`
    - We want that the pods always restart
3. Create the service
    - review the registred endpoints
4. Expose the service to your localhost
5. Test that your service is working using curl from a pod in a different ns
6. Test that all your objects can be created in a single apply command (you can use `eddygt/apphostname:1.0` has curl and some other useful tools)

Once the Jupyter container is up and running, you need to obtain the initial login token. You can do this by looking at the logs for the container

token=0195713c8e65088650fdd8b599db377b7ce6c9b10bd13766

Access to the pod via http to enter the token `http://localhost:8888/?token=<token>`

know-issue in kind there is no way like minikube to do this `minikube service nginx` for kind I access `http://localhost:8001/api/v1/namespaces/default/services/jupyter/proxy/`
but the app does not load because is looking in `http://localhost:8001/static/services/contents.js`
