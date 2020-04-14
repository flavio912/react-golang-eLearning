#!/usr/bin/env bash

COVER=0
MODULE="..."

for arg in "$@"
do
  case $arg in
    -c|--cover)
    COVER=1
    shift
    ;;        
    -m=*|--module=*)
    MODULE="${arg#*=}"
    shift
    ;;
  esac
done

# Make the container names unique
export COMPOSE_PROJECT_NAME=$(git rev-parse --short HEAD)
dc_cmd="docker-compose -f docker-compose.test.yml"

${dc_cmd} build

# start testdb if it's not running
#DOWN=0
#if [ -z `${dc_cmd} ps -q test_db` ] || [ -z `docker ps -q --no-trunc | grep $(${dc_cmd} ps -q test_db)` ]; then
#  DOWN=1
  ${dc_cmd} up -d --no-deps test_db
#fi


# Run the tests
${dc_cmd} run --rm test_api go test -v -coverprofile .testCoverage ./${MODULE}
exit_code=$?
if (($exit_code == 0 && $COVER == 1)); then
  ${dc_cmd} run --rm test_api go tool cover -func=.testCoverage
fi

# if it wasn't running stop the db
#if (($DOWN == 1)); then
  ${dc_cmd} down --volumes
#fi

exit ${exit_code}