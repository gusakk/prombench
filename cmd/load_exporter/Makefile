include ../../buildenv.mk

IMAGE_NAME              = prombench-exporter-inc
VERSION                 = latest

IMAGE_TAG               = $(IMAGE_NAME):$(VERSION)

all: build image

.PHONY: build
build: format vet staticcheck install test

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

test:
	@echo ">> running tests"
	@$(GO) test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

staticcheck: $(STATICCHECK)
	@echo ">> running staticcheck"
	@$(STATICCHECK_CMD) $(pkgs)

install:
	@echo ">> building binaries"
	@$(GO) install

image: build
	@echo ">> building docker image"
	cd $(GOPATH) && docker build --tag ${IMAGE_TAG} --force-rm -f ./src/prombench/cmd/load_exporter/Dockerfile --build-arg BINARY_PATH=./bin/load_exporter .

tag:
	@docker tag "$(IMAGE_TAG)" "$(REGISTRY_LOGIN)/$(IMAGE_NAME)"

push:
	@docker push "$(REGISTRY_LOGIN)/$(IMAGE_NAME)"

$(GOPATH)/bin/staticcheck:
	@GOOS= GOARCH= $(GO) get -u honnef.co/go/tools/cmd/staticcheck

.PHONY: all style install_dep format install test vet tarball image tag push staticcheck

# Declaring the binaries at their default locations as PHONY targets is a hack
# to ensure the latest version is downloaded on every make execution.
# If this is not desired, copy/symlink these binaries to a different path and
# set the respective environment variables.
.PHONY: $(GOPATH)/bin/staticcheck