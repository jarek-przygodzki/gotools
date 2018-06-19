.PHONY: clean build all

clean:
		find bin -type f  -delete

build:
		GOOS=darwin  go build -o bin/osx/httpl httpl.go
		GOOS=linux   go build -o bin/linux/httpl httpl.go
		GOOS=windows go build -o bin/windows/httpl.exe httpl.go

all: clean build