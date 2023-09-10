all: run

build: bin ./cmd/titleship/titleship.go
	go build -o bin/titleship ./cmd/titleship/.

install: build
	cp bin/titleship /Users/rich/bin/titleship

run: build
	bin/titleship

bin:
	mkdir -p bin
