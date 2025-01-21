default: shim op-geth

op-geth:
	go build -o op-geth "github.com/zircuit-labs/l2-geth-public/cmd/geth"
.PHONY: op-geth

shim: main.go
	go build -o shim .
