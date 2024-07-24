# bedrock-microservice-starter
Starter template for a Go based Microservice that uses AWS Bedrock 

To get started, clone this repository with:
```
git clone https://github.com/p-obrien/bedrock-microservice-starter.git
```

## Prerequiste
- Enable Claude Sonnet foundational model in your account
- Ensure you have terraform installed and configured
- Ensure your AWS Environment variables are correctly set

## Deploy EKS Cluster
Terraform code for deploying an EKS Cluster and setting up EKS Pod Identities can be found in the Infrastructure folder, for details please see the README.md file.

## Container Build

In your AWS Account create a ECR Repository with the name "bedrock-microservice-starter"

Edit the Makefile in this project to so that it contains the correct AWS Account ID and region.

Build the container with:
```
make build
```
Then push it to the container registry with
```
make push
```
Now you can deploy the application with
```
kubectl apply -f app.yaml
```

## Testing
Allow a minute or two for the load balancer to become healthy and then you can obtain the URL with:
```
kubectl get svc
```

Then call the endpoint with:
```
curl -X POST "https://your-endpoint-url.com" \
     -H "Authorization: Bearer test-api-key" \
     -F "message=say hello world like a pirate"
```

