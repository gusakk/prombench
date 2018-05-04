REGISTRY_LOGIN          = gusakk

GO                      := GO15VENDOREXPERIMENT=1 PATH=$(GOROOT)/bin:$(PATH) GOPATH=$(GOPATH) go
STATICCHECK             := $(GOPATH)/bin/staticcheck
STATICCHECK_CMD         := PATH=$(GOROOT)/bin:$(PATH) GOPATH=$(GOPATH) $(GOPATH)/bin/staticcheck
pkgs                    = $(shell $(GO) list ./... | grep -v /vendor/)