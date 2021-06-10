# Learn SQLC

## Setup SQL
1. Run `make init`
2. Edit `sqlc.yaml` to configure schema, queries, and destination folder/package
3. Run `make generate`

## Setup DB
1. Run `docker-compose up -d`
2. Open adminer at `localhost:8081`
3. Login with credentials:
    * username: `postgres`
    * password: `secret`
    * pgsql: `db:5432`
    * database: `postgres`
4. Click 'SQL Command'
5. Paste code from `schema.sql` and execute

## Demo
1. Run `go run main.go`
2. See new user in console
3. See new user in adminer
