# co2-calculator

## Build

To build the binary, either run the following go command:

```sh
$ go build -o co2-calculator main.go
```

or

```sh
$ make build
```

## Run Unit Tests


```sh
$ go test
```

or

```sh
$ make test
```

## Executing

Either use the binary:

```sh
$ ./co2-calculator --unit-of-distance=km --distance 15 --transportation-method=medium-diesel-car
```

or run through `go run main.go`:

```sh
$ go run main.go --unit-of-distance=km --distance 15 --transportation-method=medium-diesel-car
```