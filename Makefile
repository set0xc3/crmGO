.PHONY: build
build:
	templ fmt src/view
	templ generate src/view
	go build -o ./bin/program.bin src/main/main.go

.PHONY: run
run:
	templ fmt src/view
	templ generate src/view
	go run src/main/main.go || true
