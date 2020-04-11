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

To run the tests locally run:

```bash
# starts up the test database c ontainer
make test-up
# runs the test contianer
make test
```

Or you can use the shell script `api/dev_env/run_test.sh`, this is what's run in by the ci runner. See `api/middleware/*_test.go` for database intergation test examples. Fixtures are stored in `api/middleware/fixtures`. When you call `prepareDatabase()` the db is cleaned out and the fixtures inserted so you start with a fresh db for every test.
