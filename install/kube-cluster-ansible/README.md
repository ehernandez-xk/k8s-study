

The following steps install k8s cluster using Ansible scripts

https://www.digitalocean.com/community/tutorials/how-to-create-a-kubernetes-cluster-using-kubeadm-on-ubuntu-16-04

## Digital Ocean Droplets
ubuntu-16-04-x64
type: Standard 4 vCPUs 8GB

```bash
ansible-playbook -i hosts initial.yml
ansible-playbook -i hosts kube-dependencies.yml
ansible-playbook -i hosts master.yml
ansible-playbook -i hosts workers.yml
```
