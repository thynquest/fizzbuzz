default:  build

build:
	go build -mod=vendor

test:
	go test -race $(go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	go tool cover -func=coverage.out
