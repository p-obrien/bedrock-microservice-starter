build:
	docker buildx build --platform linux/arm64 ./pkg/. -f ./Dockerfile
push:
	aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin 097324129341.dkr.ecr.ap-southeast-2.amazonaws.com
	docker tag bedrock-microservice-starter:latest 097324129341.dkr.ecr.ap-southeast-2.amazonaws.com/bedrock-microservice-starter:latest
	docker push 097324129341.dkr.ecr.ap-southeast-2.amazonaws.com/bedrock-microservice-starter:latest