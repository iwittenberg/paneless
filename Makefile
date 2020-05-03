.PHONY: build install run

build: 
	go build

install:
	cp paneless.exe.manifest $(GOPATH)/bin/.
	go install -ldflags "-H windowsgui"

run:
	./paneless -preferences=$(GOPATH)/bin/preferences.json