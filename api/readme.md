# GraphQL API

Fill in your local db connection deets in `database/database.go`

Then

```
fresh
```

or

```
go run main.go
```

Send queries to localhost:8080/graphql

## Testing

To run the tests locally the shell script `api/dev_env/run_test.sh`.

```
Useage:
  ./run_test.sh --cover --module=*

Opts:
  --cover     Display a coverage report
  --module=*  Test a specific module
```

See `api/middleware/*_test.go` for database intergation test examples. Fixtures are stored in `api/middleware/fixtures`. When you call `prepareDatabase()` the db is cleaned out and the fixtures inserted so you start with a fresh db for every test.
