.PHONY: build start

build:
	go build
	chmod +x de-books

start:
	./de-books
