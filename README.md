Sandbox repo for testing argocd + kustomize deployments

Build image
```text
docker build -t go-sample-microservice:0.1 .
```

Import image into minikube registry
```text
minikube image load <image:tag>
```

Generate and apply kustomize patches for overlay
```text
kubectl apply -k .\deploy\overlays\staging
```

Forward traffic to argo server
```text
kubectl port-forward svc/argocd-server 443:443 --namespace=argocd
```

Opens a proxy to service, do not use a kubectl port-forward because it will keep the connection and won't round robin requests
```text
minikube service go-sample-microservice-service --url --namespace=example-project-staging
```
