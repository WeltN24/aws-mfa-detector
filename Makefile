PROJECT_NAME=aws-mfa-detector
VERSION=v0.4

all: bin

bin:
	GOARCH=amd64 GOOS=linux go build -o bin/linux/$(PROJECT_NAME)
	
clean:
	rm -rf bin

test:
	go test

release: clean bin
	hub release create -a "bin/linux/$(PROJECT_NAME)" -m "$(VERSION)" "$(VERSION)"