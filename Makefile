.PHONY: postgres adminer migrate migrate-down dev

docker-build:
	docker build -t sqzsvc .

docker-local-run:
	docker run -it --rm -p 5555:5555 --network host --name sqzsvc-app sqzsvc


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
