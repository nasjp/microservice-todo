.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build image
	minikube start
	eval $$(minikube -p minikube docker-env --shell=bash)
	docker build -t microservice-todo/api -f docker/Dockerfile.api .
	docker build -t microservice-todo/todo -f docker/Dockerfile.todo .

.PHONY: up
up: build ## orchestration
	kubectl cluster-info
	kubectl apply -f object/namespace.yml
	kubectl apply -f object/api.yml
	kubectl apply -f object/todo.yml


.PHONY: down
down: ## orchestration
	kubectl delete -f object/todo.yml
	kubectl delete -f object/api.yml
	kubectl delete -f object/namespace.yml
	docker rmi microservice-todo/api
	docker rmi microservice-todo/todo
	docker system prune -f
	minikube stop

.PHONY: protogen
protogen: ## gen go file from proto
	protoc -I . --go_out=plugins=grpc:. pb/*.proto

.PHONY: endpoint
endpoint: ## node port url
	minikube service api-service --url --namespace=todo
