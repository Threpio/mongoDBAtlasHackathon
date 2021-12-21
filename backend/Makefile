#MakeFile
# TODO: Make it build for all distros

dev:
	go run . start

## Build a production ready version
build: clean
	@mkdir -p dist/
	go build -o dist/backend

## Remove the release folder
clean:
	rm -rf dist/

test:
	go test ./...

fmt:
	go fmt ./...

## Build and Run a production server
server: clean build
	@mkdir -p dist/
	go build -o dist/backend
	./dist/backend start

docker-run: build
	docker build -t backend .
	docker run -p 8080:8080 -d backend