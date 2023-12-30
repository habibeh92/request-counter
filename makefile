build:
	docker build -t requester-counter:latest .

run:
	docker run --rm -p 8080:8080 -v counter-data:/app/data requester-counter:latest

test:
	go test ./...