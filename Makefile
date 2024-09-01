.PHONY: build
build:
	go build -o ./bin/program.bin cmd/myapp/myapp.go

.PHONY: run
run:
	go run cmd/myapp/myapp.go || true
