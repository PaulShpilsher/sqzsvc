.PHONY: postgres adminer migrate migrate-down dev

build:
	go build -o bin/

dev:
	go run main.go

postgres:
	docker run --rm -ti --network host \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=sqz-data \
		--name postgresql-sqz \
		postgres

adminer:
	docker run --rm -ti --network host adminer

# migrate:
# 	migrate -source file://migrations \
# 		-database postgres://postgres:secret@localhost/sqz-data?sslmode=disable up

# migrate-down:
# 	migrate -source file://migrations \
# 		-database postgres://postgres:secret@localhost/sqz-data?sslmode=disable down
