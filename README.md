# portscanner

A simple Port Scan made in Go just for fun. It's still work in progress and just scan a fixed range of ports: from 1 to 1024.

It's a simple exercise to learn Go and play with concurrency.

## Build instructions

You can build it executing:

```
go build -o dist/portscanner main.go
```

Or you can simply use the provided `Makefile`

```
make build
```

## Run instructions

You can run it executing:

```
go run main.go
```

Or you can simply use the provided `Makefile`

```
make run
```
