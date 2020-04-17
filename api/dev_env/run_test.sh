#!/usr/bin/env bash

COVER=0
BUILD=0
DEBUG=0
UNIQUE=0
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
    -d|--debug)
    DEBUG=1
    shift
    ;;
    -u|--unique)
    UNIQUE=1
    shift
    ;;
  esac
done

# Make the container names unique
if (($UNIQUE == 1)); then
export COMPOSE_PROJECT_NAME=$(git rev-parse --short HEAD)
else
export COMPOSE_PROJECT_NAME="ttc"
fi
dc_cmd="docker-compose -f docker-compose.test.yml"

# you should only need to rebuild if new modules are added
if (($KEEPALIVE == 0 || $BUILD == 1)); then
  ${dc_cmd} build
fi

${dc_cmd} up -d --no-deps test_db

printf "\n\n**********Running Tests**********\n"
# Run the tests
if (($DEBUG == 0)); then
  ${dc_cmd} run --rm --no-deps test_api go test -coverprofile .testCoverage ./${MODULE}
  exit_code=$?
else
  ${dc_cmd} run --rm --no-deps --workdir="/app/${MODULE}" test_api dlv test --api-version 2 --listen=:2345 --headless --log
  exit_code=$?
fi

if (($exit_code == 0 && $COVER == 1)); then
  ${dc_cmd} run --rm test_api go tool cover -func=.testCoverage
fi

if (($KEEPALIVE == 0)); then
  ${dc_cmd} down --volumes
fi

exit ${exit_code}