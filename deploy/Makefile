.PHONY: setup start stop clean run_script

setup:
	docker-compose up -d

start:
	docker-compose start

stop:
	docker-compose stop

clean:
	docker-compose down -v

run_script:
	./configure-dbs.sh
	./configure-replication.sh
	./create-schemas.sh

go_start:
	./env.sh
	go run ../cmd/main.go