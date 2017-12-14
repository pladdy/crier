.PHONY: config test

cover:
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out
	@echo
	@echo "'make cover html=true' to see coverage details in a browser"
ifeq ("$(html)","true")
	go tool cover -html=cover.out
endif

docs:
	godoc github.com/pladdy/lumberjack

fmt:
	go fmt -x

install:
	go get github.com/pladdy/lumberjack

test:
	go test -v -cover ./...
