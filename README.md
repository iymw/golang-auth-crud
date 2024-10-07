.env
```
PORT=3000
DB_URL="postgres://pg:pass@localhost:5432/crud"
```

1. ```docker-compose up -d```
2. ```go run migrate/migrate.go```
3. ```CompileDaemon -command="./go-crud"```