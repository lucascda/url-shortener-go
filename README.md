# Url shortener go api

### Requirements
- Docker
- Docker compose
- Go version 1.22.1+
- make
- migrate
- CompileDaemon (optional)

### How to run this project
1. Clone this repo
2. Navigate to project root folder and run `go mod tidy` to install dependencies
3. Create .env file and fill up all the variables listed there. Check .env variables docs here
4. Run `docker-compose up` to build docker services
5. Run `make migration up` to migrate database
6. Run `make run` to run api

### Enviroment variables
 - PORT: Server port Example: 3000
- JWT_SECRET: Jwt secret string Example: MyRandomJwtSecret
- DB_USER: Postgresql user Example: root
- DB_PASSWORD: Postgresql password Example: mypass1
- DB_NAME: Postgresql database name Example: url_api
- DB_PORT: Host port to attach postgres docker db Example: 5432


