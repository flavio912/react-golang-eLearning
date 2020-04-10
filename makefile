test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes

test-db-up:
	docker-compose -f docker-compose.test.yml up -d --build test_db 

test-db-down:
	docker-compose -f docker-compose.test.yml down --volumes test_db