service := {{<service_name>}}
version := 0.0.0
docker_org := {{<docker_registry>}}
gcloud_proj := {{<project_id>}}
cluster := {{<cluster_name>}}
docker-image := ${docker_org}/${service}:${version}
root := $(abspath $(shell pwd))
port := {{<port>}}

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
	sudo docker build -t $(docker-image) .

docker-dev:
	sudo make docker-build
	sudo make docker-run

docker-push:
	sudo docker push $(docker-image)

docker-run:
	@docker run -itp $(port):$(port)  $(docker-image)

bumpversion-patch:
	bumpversion patch --allow-dirty

bumpversion-minor:
	bumpversion minor --allow-dirty

bumpversion-major:
	bumpversion major --allow-dirty

bootstrap-deploy:
	gcloud container clusters get-credentials ${cluster} --zone us-central1-c --project ${gcloud_proj}
	make docker-build
	make docker-push
	kubectl create -f deployments/k8s/service.yml
	kubectl create -f deployments/k8s/deploy.yml

deploy:
	git add .
	git commit -m "${version}"
	git push origin master
	gcloud container clusters get-credentials ${cluster} --zone us-central1-c --project ${gcloud_proj}
	sudo make docker-build
	sudo make docker-push
	kubectl apply -f deployments/k8s/deploy.yml

ci-deploy:
	./google-cloud-sdk/bin/gcloud container clusters get-credentials ${cluster} --zone us-central1-c --project ${gcloud_proj}
	./kubectl apply -f deployments/k8s/deploy.yml

purge:
	go clean
	rm -rf $(root)/vendor

test-workstation:
	go test ./... --env=workstation -v 3
