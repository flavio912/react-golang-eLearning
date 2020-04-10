#!/usr/bin/env bash
set -x

# Make the container names unique
export COMPOSE_PROJECT_NAME=$(git rev-parse --short HEAD)
docker_compose_cmd="docker-compose -f docker-compose.test.yml"
${docker_compose_cmd} build

${docker_compose_cmd} up -d --no-deps test_db

# Run the tests
${docker_compose_cmd} run --rm test_api
exit_code=$?

# Clean up
${docker_compose_cmd} down --volumes

exit ${exit_code}