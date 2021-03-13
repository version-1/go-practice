

## How to run


1. run mysql
```
docker run  -v -e MYSQL_ROOT_PASSWORD -e MYSQL_ALLOW_EMPTY_PASSWORD=true -p 3306:3306 mysql

```

2. db setup

```
go run cmd/rest-api/migrate/main.go
go run cmd/rest-api/seed/main.go
```

3.  run

```
go run cmd/rest-api/main.go
```
