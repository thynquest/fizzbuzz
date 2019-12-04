default:  build

build:
	go build -mod=vendor

test:
	go test ./... -cover
