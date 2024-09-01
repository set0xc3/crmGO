.PHONY: build
build:
	templ fmt internal/view
	templ generate internal/view
	go build -o ./bin/program.bin cmd/myapp/myapp.go

.PHONY: run
run:
	templ fmt internal/view
	templ generate internal/view
	go run cmd/myapp/myapp.go || true
