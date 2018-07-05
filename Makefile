.PHONY: clean build all

GIT_COMMIT=$(shell git rev-parse HEAD)

clean:
		find bin -type f  -delete
build:
		GOOS=darwin  go build -o bin/osx/goexec -ldflags "-X main.GitCommit=$(GIT_COMMIT)" goexec.go
		GOOS=linux   go build -o bin/linux/goexec -ldflags "-X main.GitCommit=$(GIT_COMMIT)" goexec.go
		GOOS=windows go build -o bin/windows/goexec.exe -ldflags "-X main.GitCommit=$(GIT_COMMIT)" goexec.go

all: clean build