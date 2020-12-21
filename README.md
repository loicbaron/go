# go

Learning GoLang.

# Tour of Go

Some examples from <a href="https://tour.golang.org/" target="_blank">https://tour.golang.org/</a>

# import within the same package

```
cd src/samePackage
go run main.go state.go
go build && ./samePackage
```

# hello package

This is my first go package.

## Setup environment

```
cd src/hello
export GOPATH=$GOPATH:$(pwd)
```

## Run

```
go run main.go
go build && ./hello
```
