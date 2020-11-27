build:
	GOOS=linux GOARCH=amd64 go build -o br-seed-go
	docker build -t br-seed-go .

run:
	docker run -p 50051:50051 br-seed-go


