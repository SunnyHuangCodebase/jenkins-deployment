# Deploy Jenkins on Minikube
Setup Jenkins on a Minikube Kubernetes cluster using Pulumi's Infrastructure as Code written in Go or  `kubectl` commands.  

## Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Pulumi](https://www.pulumi.com/docs/install/) (optional, or use kubectl alone) 

### Change To Working Directory
Pulumi:
```sh
cd /path/to/this/project/jenkins-deployment
```

Kubectl:
```sh
cd /path/to/this/project/jenkins-deployment/manifests
``` 

### Create/Update Resources
Pulumi
```sh
pulumi up --yes --skip-preview
```

Kubectl:
```sh
kubectl apply -f jenkins.yaml
```

### Forward Jenkins Service Port
```sh
kubectl port-forward `kubectl get service -o name | grep jenkins` 8000:8000
```

### Access Jenkins:
[http://localhost:8000](http://localhost:8000)

### Delete Resources
Pulumi:
```sh
pulumi destroy --skip-preview --yes
```

Kubectl:
```sh
kubectl delete -f jenkins.yaml
```
