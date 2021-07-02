#!/usr/bin/make -f

export CGO_ENABLED=0

all: test build

# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out

# Perform a test package
dry-run:
	goreleaser --snapshot --skip-publish --rm-dist

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	controller-gen crd:trivialVersions=true rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Run golint against code
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Generate code
generate:
	controller-gen object:headerFile=./hack/boilerplate.go.txt paths="./..."
	client-gen --input-base=github.com/skpr/aws-auth-mapper/apis \
               --input="iamauthenticator/v1beta1" \
               --go-header-file=./hack/boilerplate.go.txt \
               --output-package=github.com/skpr/aws-auth-mapper/internal/ \
               --clientset-name=clientset

.PHONY: *
