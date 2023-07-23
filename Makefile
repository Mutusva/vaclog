.PHONY: docs
docs:
	swag init --parseDependency --parseDepth 1 -g cmd/vaclog/main.go --output docs/swag -q

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -cover -coverprofile=coverage.out -race ./...


.PHONY: build
build:
	go build  -o ./build/main ./cmd/vaclog/main.go

.PHONY: build-image
build-image:
	docker build vaclog .

.PHONY: run
run:
	go run ./cmd/vaclog/main