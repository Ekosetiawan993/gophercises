## Hello world of rest API using Go 1.22.2

It's me again try to learn Go.

To use this simple program, simply run it using

```bash
go run main.go
```

or compile it then run the binary

```bash
go build main.go
main.exe
```

#### To hit the API

```bash
# index api
curl http://localhost:8089
# get hello api
curl http://localhost:8089/hello
# post hello api
curl -X POST http://localhost:8089/hello
# get hello with id api
curl http://localhost:8089/hello/100
```
