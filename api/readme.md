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
Usage:
  ./run_test.sh

Opts:
  --cover       Display a coverage report
  --module=*    Test a specific module (default is ..., ie all modules in api)
  --build       Rebuild (default false)
  --keep-alive  Doesn't stop the container
  --debug       Start in debug mode
  --html        Creates a coverage report at api/coverage_report.html
```

Any other args will get passed straight to go test. e.g. `./api/dev_env/run_test.sh -v --run Managers` will run all tests which match the Managers regex in verbose mode

See `api/middleware/*_test.go` for database intergation test examples. Fixtures are stored in `api/middleware/fixtures`. When you call `prepareTestDatabase()` the db is cleaned out and the fixtures inserted so you start with a fresh db for every test.

### Debugging tests

You can only debug one module at a time (limitation of Delve). Run `./run_test.sh --debug --module=***`, then launch the debug config from vscode (or connect to it using `dlv connect host:2345` to use the dlv cli).

N.B. you might need to change the host ip in `launch.json`. To find out what yours is start the debug server then run: `docker inspect ttc_test_api | grep IPAddress`. The ip will change with each new network that's used :shrug:
