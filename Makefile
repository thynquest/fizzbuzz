GO := go

## BUILD ##
.PHONY: build
build:
	$(GO) build cmd/app/main.go
	@rm main

## DEPENDENCIES ##
.PHONY: tidy
tidy:
	@$(GO) mod tidy

.PHONY: unittest
unittest:
	$(GO) test ./...  -coverprofile=coverage.out.tmp
	cat coverage.out.tmp  > coverage.out
	rm coverage.out.tmp


.PHONY: coverage
coverage: unittest
	$(GO) tool cover -func coverage.out
	
## RUN ##
.PHONY: run
run:
	@export ENV_FILE="app.env" && \
	$(GO) run cmd/app/main.go
