# Start ping subnets:

```bash
Usage:
    go-ping-list [-s start IP in subnet] [-e end IP in subnet]

Example:
    go-ping-list -s 172.20.13.1 -e 172.20.13.254
```

ICMP Ping library for Go, inspired by
[go-fastping](https://github.com/tatsushid/go-fastping)

For a full ping example, see
[cmd/ping/ping.go](https://github.com/sparrc/go-ping/blob/master/cmd/ping/ping.go)


# Installation:

Install with "dep" package:

```bash
go get github.com/cp38510/go-ping-list
cd $GOPATH/src/github.com/cp38510/go-ping-list
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
$GOPATH/bin/dep ensure
go build go-ping-list.go
```

Or hand install:

```bash
go get github.com/sparrc/go-ping
go get github.com/cp38510/go-ping-list
go build $GOPATH/src/github.com/cp38510/go-ping-list/go-ping-list.go
#$GOPATH/bin/go-ping-list
```
