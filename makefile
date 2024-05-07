.PHONY: run clean
.PHONY: migration up
.PHONY: migration down

run:
	CompileDaemon -directory="./src" -build="go build -o api" -command="./src/api"

clean:
	rm api

build:
	go build -o api
migration up:
	migrate -database postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable \
	-path ./src/migrations up

migration down:
	migrate -database postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable \
	-path ./src/migrations down