build:
	docker buildx build --platform linux/arm64 ./pkg/. -f ./Dockerfile -t bedrock-microservice-starter
push:
	aws ecr get-login-password --region <region> | docker login --username AWS --password-stdin <aws account number>.dkr.ecr.<region>.amazonaws.com
	docker tag bedrock-microservice-starter:latest <aws account number>.dkr.ecr.<region>.amazonaws.com/bedrock-microservice-starter:latest
	docker push <aws account number>.dkr.ecr.<region>.amazonaws.com/bedrock-microservice-starter:latest