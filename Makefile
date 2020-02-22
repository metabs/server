NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

BINARY_NAME?=app
DIR_OUT=$(CURDIR)/out

.PHONY: all clean deps build install

all: clean build install

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

test-integration: tools.format tools.vet
	@command -v godog >/dev/null ; if [ $$? -ne 0 ]; then \
			echo "--> installing godog"; \
	go get github.com/cucumber/godog/cmd/godog; \
	fi

	@printf "$(OK_COLOR)==> Spinning up docker-compose$(NO_COLOR)\n"
	@docker-compose up -d

	@printf "$(OK_COLOR)==> Running integration tests$(NO_COLOR)\n"
	go test ./tests -godog -stop-on-failure -v

test-unit: tools.format tools.vet
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
#-- tools
#---------------
.PHONY: tools tools.errcheck tools.golint tools.goimports tools.format tools.vet
tools: tools.errcheck tools.goimports tools.format tools.lint tools.vet

tools.goimports:
	@command -v goimports >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing goimports"; \
		go get golang.org/x/tools/cmd/goimports; \
	fi
	@echo "$(OK_COLOR)==> checking imports 'goimports' tool$(NO_COLOR)"
	@goimports -l -w cmd internal &>/dev/null | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

tools.format:
	@echo "$(OK_COLOR)==> formatting code with 'gofmt' tool$(NO_COLOR)"
	@gofmt -l -s -w cmd internal | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

tools.lint:
	@command -v golint >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing golint"; \
		go get github.com/golang/lint/golint; \
	fi
	@echo "$(OK_COLOR)==> checking code style with 'golint' tool$(NO_COLOR)"
	go list ./... | xargs -n 1 golint -set_exit_status

tools.vet:
	@echo "$(OK_COLOR)==> checking code correctness with 'go vet' tool$(NO_COLOR)"
	go vet ./...

tools.errcheck:
	@command -v errcheck >/dev/null ; if [ $$? -ne 0 ]; then \
			echo "--> installing errcheck"; \
			go get -u github.com/kisielk/errcheck; \
		fi
	@echo "$(OK_COLOR)==> checking proper error handling with 'go errcheck' tool$(NO_COLOR)"
	@errcheck -ignoretests ./cmd/... ./internal/... ./pkg/...
