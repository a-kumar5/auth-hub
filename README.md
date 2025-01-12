# auth-hub

## Building the application

```bash
go build -o auth-hub ./cmd/main/main.go

## to run the application
./auth-hub
```

## application db setup using docker

```bash
docker run --name postgres -e POSTGRES_PASSWORD=admin -p 5432:5432 -d postgres
docker exec -it postgres createdb -U postgres authdb
```