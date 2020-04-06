.PHONY: build

build:
	go build

exec:
	$(FENNEL) main.fnl

server:
	$(FENNEL) server.fnl

client:
	$(FENNEL) client.fnl
