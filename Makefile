.PHONY: clean build all

clean:
		find bin -type f  -delete

build:
		GOOS=darwin  go build -o bin/osx/tcpl tcpl.go
		GOOS=linux   go build -o bin/linux/tcpl tcpl.go
		GOOS=windows go build -o bin/windows/tcpl.exe tcpl.go

all: clean build