requirements_freeze:
	pip freeze > requirements.txt

requirements_install:
	pip install -r requirements.txt

venv_create:
	python3 -m venv .venv

venv_activate:
	source .venv/bin/activate

nix-shell:
	nix-shell ./shell.nix

grpc_python:
	python -m grpc_tools.protoc -I ../protobufs/ --python_out=./ --grpc_python_out=./k ../protobufs/query.proto
