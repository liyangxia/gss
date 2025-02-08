# Go Simple Server (GSS)

A simple HTTP file server written in Go.

## Features

- Serves files from the current directory
- Logs request details
- Disables caching for served files
- Configurable port and LAN access

## Install

To install the project using `go install`, run the following command:

```sh
go install github.com/liyangxia/gss/src/@latest
```

This will install the `gss` binary in your `$GOPATH/bin` directory.

## Build

To build the project manually, ensure you have Go installed and run the following commands:

```sh
git clone https://github.com/liyangxia/gss.git
cd gss
make build
```

This will create a binary in the `dist` directory.

## Usage

To run the server, use the following command:

```sh
./dist/gss [--port PORT] [--lan]
```

### Options

- `--port PORT`: Specify the port to serve on (default is 8080)
- `--lan`: Allow LAN access (default is false)

### Example

Serve files on port 8080:

```sh
./dist/gss --port 8080
```

Serve files on port 8080 and allow LAN access:

```sh
./dist/gss --port 8080 --lan
```

## Contributing

PRs accepted. Please ensure your code adheres to the project's coding standards.

## License

MIT © liyangxia
