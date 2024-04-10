.PHONY: run

run:
	CompileDaemon -directory="./src" -build="go build -o api" -command="./src/api"