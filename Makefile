.PHONY: clean build all

clean:
		find bin -type f  -delete

build:
		GOOS=darwin  go build -o bin/osx/goexec -ldflags "-X main.GitCommit=$(git rev-list -1 HEAD)" goexec.go
		GOOS=linux   go build -o bin/linux/goexec -ldflags "-X main.GitCommit=$(git rev-list -1 HEAD)" goexec.go
		GOOS=windows go build -o bin/windows/goexec.exe -ldflags "-X main.GitCommit=$(git rev-list -1 HEAD)" goexec.go

all: clean build