service := {{<service_name>}}
docker-image := {{<docker_registry>}}/{{<service_name>}}:0.0.0
root := $(abspath $(shell pwd))
port := 7001

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bootstrap:
	go mod init $(service)
	make init

init:
	go mod tidy

dev:
	go run main.go

docker-build:
	docker build -t $(docker-image) .

docker-dev:
	make docker-build
	make docker-run

docker-push:
	docker push $(docker-image)

docker-run:
	@docker run -itp $(port):$(port)  $(docker-image)

bumpversion-patch:
	bumpversion patch --allow-dirty

bootstrap-deploy:
	gcloud container clusters get-credentials ${cluster} --zone us-central1-c --project ${gcloud_proj}
	make docker-build
	make docker-push
	kubectl create -f k8s/service.yml
	kubectl create -f k8s/deploy.yml

deploy:
	git add .
	git commit -m "${version}"
	git push origin master
	gcloud container clusters get-credentials ${cluster} --zone us-central1-c --project ${gcloud_proj}
	make docker-build
	make docker-push
	kubectl apply -f k8s/deploy.yml

purge:
	go clean
	rm -rf $(root)/vendor
