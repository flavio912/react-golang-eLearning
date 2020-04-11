test:
	docker-compose -f docker-compose.test.yml build test_api
	docker-compose -f docker-compose.test.yml run --rm test_api

test-up:
	docker-compose -f docker-compose.test.yml up -d --build test_db

test-down:
	docker-compose -f docker-compose.test.yml down --volumes