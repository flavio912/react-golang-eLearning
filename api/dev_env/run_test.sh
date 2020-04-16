#!/usr/bin/env bash

COVER=0
BUILD=0
KEEPALIVE=0
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
    -k|--keep-alive)
    KEEPALIVE=1
    shift
    ;;
    -b|--build)
    BUILD=1
    shift
    ;;
  esac
done

# Make the container names unique
export COMPOSE_PROJECT_NAME=$(git rev-parse --short HEAD)
dc_cmd="docker-compose -f docker-compose.test.yml"

# you should only need to rebuild if new modules are added
if (($KEEPALIVE == 0)); then
  ${dc_cmd} build
fi

# start testdb if it's not running
#DOWN=0
#if [ -z `${dc_cmd} ps -q test_db` ] || [ -z `docker ps -q --no-trunc | grep $(${dc_cmd} ps -q test_db)` ]; then
#  DOWN=1
  ${dc_cmd} up -d --no-deps test_db
#fi


# Run the tests
${dc_cmd} run --rm --no-deps test_api go test -coverprofile .testCoverage ./${MODULE}
exit_code=$?
if (($exit_code == 0 && $COVER == 1)); then
  ${dc_cmd} run --rm test_api go tool cover -func=.testCoverage
fi

if (($KEEPALIVE == 0)); then
  ${dc_cmd} down --volumes
fi

exit ${exit_code}