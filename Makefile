.PHONY: all
all: build
FORCE: ;

SHELL  := env GETIR_GO_CHALLANGE_ENV=$(GETIR_GO_CHALLANGE_ENV) $(SHELL)
GETIR_GO_CHALLANGE_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

# TODO: Include env variables
#ifneq (,$(wildcard ./.env))
#    include .env
#    export
#    ENV_FILE_PARAM = --env-file .env
#endif
build: docker

build-api:
	go build -tags $(GETIR_GO_CHALLANGE_ENV) -o ./bin/api api/server.go

#build-cmd:
#	go build -tags $(GETIR_GO_CHALLANGE_ENV) -o ./bin/cmd cmd/main.go

docker:
	docker-compose -f docker-compose-api.yml up -d

#linux-binaries:
#	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(GETIR_GO_CHALLANGE_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/api api/server.go
#	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(GETIR_GO_CHALLANGE_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/search cmd/main.go

ci: dependencies test

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/record/interface.go -destination=usecase/record/mock/record.go -package=mock
	@~/go/bin/mockgen -source=usecase/configuration/interface.go -destination=usecase/configuration/mock/configuration.go -package=mock

test:
	go test -tags testing ./...

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done