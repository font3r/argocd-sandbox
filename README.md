# Build image
docker build -t go-sample-microservice:0.1 .

# Import image into minikube registry
minikube image load <image:tag>

# Generate and apply kustomize patches 
kubectl apply -k .\deploy\overlays\staging

# Opens a proxy to service
# do not use a kubectl port-forward because it will keep the connection and won't round robin requests
minikube service go-sample-microservice-service --url --namespace=example-project-staging