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


To install the native Go ping list executable:

```bash
go get github.com/sparrc/go-ping
go get gitlab.techlab.s7.ru/a.s.reshetnikov/go-ping-list
go build $GOPATH/src/gitlab.techlab.s7.ru/a.s.reshetnikov/go-ping-list/go-ping-list.go
#$GOPATH/bin/go-ping-list
```
