# Movies List Crud App

A basic HTTP CRUD app with basic models and table joins.

## Building

This projects uses `sqlite3` for the database so make sure that's installed.

There is a `db/seeds.sh` script that will seed the database with data which also truncate the tables and reset ID auto increment count.

Run `go run main.go` to run the app and start the HTTP server on `http://localhost:8080`.
