kubectl get componentstatuses

kubectl logs --previous nginx-65899c769f-2pgzk

kubectl describe replicationcontroller

kubectl logs -l app=nginx --all-containers=true

repasar deployment rolling update and recreate

kubectl get secrets --all-namespaces
work with kubeadm

kubectl drain k8s-worker04 --ignore-daemonsets


3 hours
24 questions
74%


19% Core Concepts, API, master and nodes componets
 - https://kubernetes.io/docs/concepts/overview/components/
12% Installation, configuration and Validation - Design and Build
 - https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/cluster-interactive/
12% security - RBAC, security contexts, TLS, network and pod security policies
 - https://kubernetes.io/docs/reference/access-authn-authz/controlling-access/
11% Cluster maintenance - Upgrade nodes by drain/cordon, etcd maintenance
 - https://kubernetes.io/docs/tasks/administer-cluster/cluster-management/
11% Networking, Services, ingress, cluster and pod netwoking
 - https://kubernetes.io/docs/concepts/services-networking/service/
10% Troubleshooting - cluster component, application, networking failure
 - https://kubernetes.io/docs/tasks/debug-application-cluster/debug-cluster/
8% Application Life Cycle Management- rollout, rollback, update and scale applications
 - https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
7% Storage - Storage objects, configure applications with storage
 - https://kubernetes.io/docs/tasks/configure-pod-container/configure-volume-storage/
5% Logging & Monitoring - node and application health, cluster component
 - https://kubernetes.io/docs/concepts/cluster-administration/logging/
5% Scheduling, Labels, resources constraints, scheduler configuraiton
 - https://kubernetes.io/docs/concepts/configuration/assign-pod-node/



https://kubernetes.io/docs/reference/kubectl/cheatsheet/

blogs
https://kubedex.com/7-5-tips-to-help-you-ace-the-certified-kubernetes-administrator-cka-exam/
https://blog.heptio.com/how-heptio-engineers-ace-the-certified-kubernetes-administrator-exam-93d20af32557
https://github.com/kubernauts/Kubernetes-Learning-Resources

Courses
LFS258 Linux Foundation - Kubernetes Fundamentals https://training.linuxfoundation.org/training/kubernetes-fundamentals/

Books
Managing Kubernetes (Brendan Burns and Craig Tracey)
Kubernetes Up and Running (Kelsey Hightowe, Brandan Burns and Joe Beda)

command line tips
vi 
g end to the page
gg begining to the page
insert line
delete line
end to line
begning line


alias
alias kgp="kubectl get pods"

use help
create object with run
kubecetl run -h

some tools to troubleshoot
dig
host

kubectl api-resources

https://www.cncf.io/certification/cka/


see for some videos
https://github.com/walidshaari/Kubernetes-Certified-Administrator

I did this for auto completion
https://medium.com/merapar/fixing-bash-autocompletion-on-macos-for-kubectl-and-kops-e87f019652e8
