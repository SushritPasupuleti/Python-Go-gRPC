# Python Go gRPC

A Python gRPC service that is accessed by a REST API built with Go-chi.

## Layout

- `snek/` - Python gRPC service
- `server/` - Go-chi REST API

## Setup

For NixOS users, you can use the provided `shell.nix` to get a development environment.

```bash
nix-shell
```

Refer to the `Makefile`s in each directory for available commands.

## Usage

Regenerating gRPC code:

```bash
make generate
```

Running the Python gRPC service:

```bash
cd snek
make run
```

Running the Go-chi REST API:

```bash
cd server
make run # or make dev
```
