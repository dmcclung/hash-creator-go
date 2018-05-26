# hash-creator-go
## Setup
Using go1.10.1 windows/amd64

`export GOPATH=<PATH to hash-creator-go cloned repo>`

`go build github.com/hash-creator-go/hashstore`

`go build github.com/hash-creator-go/hashstore-rest-app`

## Tests
`go test github.com/hash-creator-go/hashstore`

`go test github.com/hash-creator-go/hashstore-rest-app`

`go test -bench=BenchmarkHashStore github.com/hash-creator-go/hashstore-rest-app`

## Running
`go install github.com/hash-creator-go/hashstore`

`go install github.com/hash-creator-go/hashstore-rest-app`

`$GOPATH/bin/hashstore-rest-app.exe [-port=<port, default is 8080>]`
