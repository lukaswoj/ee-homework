image_name = lukaswoj/ee-homework

minikube-start:
	minikube start --addons=ingress
apply-k8s-resources:
	kubectl apply -f kubernetes/namespace.yaml
	kubectl apply -f kubernetes/deployment.yaml
	kubectl apply -f kubernetes/service.yaml
	kubectl apply -f kubernetes/ingress.yaml
image-build:
	docker build -t ${image_name} ./app
image-push:
	docker push ${image_name}
restart-pods:
	kubectl -nee-homework-dev rollout restart deployment app
