APP_NAME ?= demo-service
IMAGE_NAME ?= $(APP_NAME)
IMAGE_TAG ?= local
DOCKERFILE ?= app/Dockerfile
BUILD_CONTEXT ?= app

.PHONY: test docker-build docker-run terraform-fmt

test:
	cd app && go test ./...

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f $(DOCKERFILE) $(BUILD_CONTEXT)

docker-run:
	docker run --rm -p 8080:8080 -e APP_VERSION=$(IMAGE_TAG) $(IMAGE_NAME):$(IMAGE_TAG)

terraform-fmt:
	terraform fmt -recursive terraform
