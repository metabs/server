NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

BINARY_NAME?=app
DIR_OUT=$(CURDIR)/dist

.PHONY: all clean deps build

all: clean build

clean:
	@printf "$(OK_COLOR)==> Cleaning project$(NO_COLOR)\n"
	rm -f ${DIR_OUT}/${BINARY_NAME}

build:
	@printf "$(OK_COLOR)==> Building server binary$(NO_COLOR)\n"
	go build -o ${DIR_OUT}/${BINARY_NAME} ./cmd/server/


#---------------
#-- tests
#---------------
.PHONY: tests test-integration test-unit
tests: test-integration test-unit

test-integration:
	@printf "$(OK_COLOR)==> Spinning up docker-compose$(NO_COLOR)\n"
	@docker-compose up -d

	@printf "$(OK_COLOR)==> Running integration tests$(NO_COLOR)\n"
	go test ./tests -cucumber -stop-on-failure -v

test-unit:
	@printf "$(OK_COLOR)==> Unit Testing$(NO_COLOR)\n"
	go test -v -race -cover ./...

#---------------
#-- deps
#---------------
.PHONY: deps
deps:
	@printf "$(OK_COLOR)==> Installing deps$(NO_COLOR)\n"
	go mod tidy
	go mod vendor

#---------------
#-- lint
#---------------
.PHONY: lint
lint:
	which golangci-lint; if [ $$? -ne 0 ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.21.0; fi
	golangci-lint run --modules-download-mode vendor --fix