all: run

build: bin ./cmd/titleship/titleship.go
	go build -o bin/titleship ./cmd/titleship/titleship.go

install: build
	cp bin/titleship /Users/rich/bin/titleship

run: build
	bin/titleship

bin:
	mkdir -p bin