grpc_python:
	@echo "Generating python grpc code"
	@python -m grpc_tools.protoc -I protobufs/ --python_out=./snek --grpc_python_out=./snek protobufs/*.proto
	@echo "Finished generating python grpc code"

grpc_go:
	@echo "Generating go grpc code"
	@protoc --proto_path=protobufs protobufs/*.proto --go_out=server --go-grpc_out=server
	@echo "Finished generating go grpc code"

generate: grpc_python grpc_go

nix-shell:
	@echo "Entering nix-shell"
	@nix-shell ./shell.nix --run fish
