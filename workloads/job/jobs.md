## Jobs

It is a Job controller for managing Pods. These pods generally run until successful completion. The Job object coordinates running a number of pods in parallel.
If POD fails will create a new POD

### Oneshot
Get the basic yaml and add the command for the container

`job.spec.template.spec.containers.command`

```bash
# this job does not do anything
kubectl create job my-job --image=busybox --dry-run -o yaml > kuard-simple-job.yaml
# cleanup and add the container command
kubectl apply -f kuard-simple-job.yaml
kubectl logs my-job-rg2m8
hola
```

You can pass the command after the `--` for example

```bash
# this does the same as the above example
kubectl create job my-job --image=busybox --dry-run -o yaml -- echo hola > kuard-simple-job.yaml
```

`job.spec.template.spec.restartPolicy` options: Always, OnFailure, Never

Never: will not restart instead will create more pods (be careful)

### POD FAILURE

Create a pod that fails

```bash
kubectl create job my-job --image=busybox --dry-run -o yaml -- ls non-file > kuard-fail-job.yaml
kubectl apply -f kuard-fail-job.yaml
# see the failing pods
kubectl get pods -w
```

## Parallelism

```bash
kubectl create job my-job --image=alpine --dry-run -o yaml -- sleep 5 > kuard-parallel-job.yaml
# edit the yaml
kubectl apply -f kuard-parallel-job.yaml
```

`job.spec.completions` Specifies the desired number of successfully finished pods
if the completations is not set the job will be set as pool worker
`job.spec.parallelism` how many pods will run in parallel