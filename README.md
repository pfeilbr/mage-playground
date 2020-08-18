# mage-playground

learn [mage](https://magefile.org/), a make/rake-like build tool using Go. You write plain-old go functions, and Mage automatically uses them as Makefile-like runnable targets.

## Install

```sh
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```


## Demo

```sh
git clone https://github.com/pfeilbr/mage-playground
cd mage-playground
go mod init mage-playground
# NOTE: `go.mod` must exist in same directory as `magefile.go`

# show targets
mage

# run "hello" target
mage hello
```