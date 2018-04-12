linux:
	CGO_ENABLED=0 GOOS=linux go build -v

build:
	go build -v

run: build
	./ok-tcp-server --host localhost:1234
