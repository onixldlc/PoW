# PoW (Packet Over WebSocket)

This project is a Go CLI application that facilitates communication over WebSocket using a proxy architecture. It allows users to send and receive messages through a middleware layer.

## Features

- **Sender**: Sends packet to the middleware.
- **Receiver**: Listens for incoming messages from a middleware.
- **Middleware**: Acts as a connector between the sender and receiver, processing messages as needed.

## Installation

To install the project, clone the repository and navigate to the project directory:

```bash
git clone https://github.com/onixldlc/PoW.git
cd packet-over-websocket
```

Then, build the project using Go:

```bash
go build -o build/PoW ./
```

## Usage

You can run the CLI program with the following commands:

- `sender` or `s`: Run the sender functionality.
- `receiver` or `r`: Run the receiver functionality.
- `middleware` or `m`: Run the middleware functionality.

Options:
- `<host>`: ip to listen to
- `-p`: setup listener port 

### Examples

To run the sender:

```bash
./PoW s
```

To run the receiver:

```bash
./PoW r -p 8081
```

To run the middleware:

```bash
./PoW m 172.0.0.0
./PoW m 172.0.0.0 -p 1111
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.