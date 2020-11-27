# Hello Service

This is the Seed go Hello service

Generated with

```
micro new br-seed-go
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

.PHONY: build

```
go build -o hello *.go
```

.PHONY: test

```
cd tests
go test -v 
```

.PHONY: docker

```
docker build . -t hello:latest
```

.PHONY: client

```
cd client
go run main.go
```
