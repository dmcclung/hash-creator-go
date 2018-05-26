# jcloud-assignment
## Setup
Using go1.10.1 windows/amd64

`export GOPATH=<PATH to jcloud-assignment cloned repo>`

`go build github.com/dylan-jcloud-assignment/hashstore`

`go build github.com/dylan-jcloud-assignment/hashstore-rest-app`

## Tests
`go test github.com/dylan-jcloud-assignment/hashstore`

`go test github.com/dylan-jcloud-assignment/hashstore-rest-app`

`go test -bench=BenchmarkHashStore github.com/dylan-jcloud-assignment/hashstore-rest-app`

## Running
`go install github.com/dylan-jcloud-assignment/hashstore`

`go install github.com/dylan-jcloud-assignment/hashstore-rest-app`

`$GOPATH/bin/hashstore-rest-app.exe [-port=<port, default is 8080>]`
