.PHONY: build
build:
	go build -o inven-store .

.PHONY: docker
docker:
	docker build -t inven-store .

.PHONY: run 
run:
	docker run --name inven-store -d -p 8080:8080 inven-store 

.PHONY: tag
tag:
	docker tag inven-store:latest renegmedal/inven-store:1.0.1

.PHONY: push
push:
	docker push renegmedal/inven-store:1.0.1

.PHONY: up
up:
	docker-compose up --build -d 
