build:
	go build -C cmd/manager/ -o ../../bin/

run: build
	./bin/manager -env ./.env