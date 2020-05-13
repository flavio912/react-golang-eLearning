#!/usr/bin/env bash

COVER=0
BUILD=0
DEBUG=0
UNIQUE=0
HTML=0
KEEPALIVE=0
MODULE="..."
testApiName='ttc_test_api'

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
    --html)
    HTML=1
    shift
    ;;
  esac
done

# Make the container names unique
if [ $UNIQUE == 1 ]; then
export COMPOSE_PROJECT_NAME=$(git rev-parse --short HEAD)
else
export COMPOSE_PROJECT_NAME="ttc"
fi
dc_cmd="docker-compose -f docker-compose.test.yml"

# you should only need to rebuild if new modules are added
if [ $KEEPALIVE == 0 ] || [ $BUILD == 1 ]; then
  if [ "$(docker ps -aq -f name=$testApiName)" ]; then
    printf "\n\n**********Removing Old Test Containers**********\n"
    docker stop $testApiName
    docker rm $testApiName
  fi

  printf "\n**********Building Test Containers**********\n"
  $dc_cmd build
fi

$dc_cmd up -d --no-deps test_db

run_cmd() {
  exit_code=0
  if [ $KEEPALIVE == 1 ]; then
    if ! docker ps --format '{{.Names}}' | grep -w $testApiName &> /dev/null; then
      if [ "$(docker ps -aq -f status=exited -f name=$testApiName)" ]; then
        docker rm $testApiName
      fi
      set -x
      $dc_cmd run -d --rm --no-deps --name $testApiName test_api tail -f /dev/null
      set +x
    fi

    if [ -n "$2" ]; then #workdir
      set -x
      docker exec --workdir="$2" ttc_test_api $1
      exit_code=$?
      set +x
    else
      set -x
      docker exec ttc_test_api $1
      exit_code=$?
      set +x
    fi
  else
    if [ -n "$2" ]; then #workdir
      $dc_cmd run --rm --no-deps --workdir="$2" test_api $1
      exit_code=$?
    else
      $dc_cmd run --rm --no-deps test_api $1
      exit_code=$?
    fi
  fi

  return $exit_code
}

printf "\n**********Running Tests**********\n"

if [ $DEBUG == 0 ]; then
  run_cmd "go test -coverprofile .testCoverage ./${MODULE} -p=1 $*"
  exit_code=$?
else
  run_cmd "dlv test --api-version 2 --listen=:2345 --headless --log -- $*" "/app/${MODULE}" 
  exit_code=$?
fi

if [ $exit_code == 0 ] && [ $COVER == 1 ]; then
  run_cmd "go tool cover -func=.testCoverage"
fi
if [ $exit_code == 0 ] && [ $HTML == 1 ]; then
  run_cmd "go tool cover -html=.testCoverage -o=cover_report.html"
fi

if [ $KEEPALIVE == 0 ]; then
  $dc_cmd down --volumes
fi

exit ${exit_code}