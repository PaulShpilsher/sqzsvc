
.PHONY: docker-build
docker-build:
	docker build -t sqzsvc .

.PHONY: docker-run
docker-run:
	docker run -it --rm -p 5555:5555 --network host --name sqzsvc-app sqzsvc


.PHONY: build
build:
	go build -o bin/

.PHONY: serve
serve:
	go run main.go

.PHONY: dev
dev:
	go run main.go -debug=true -port=5555 -db-dns="localhost user=postgres password=secret dbname=sqz-data port=5432 sslmode=disable"

.PHONY: postgres
postgres:
	docker run --rm -ti --network host \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=sqz-data \
		--name postgresql-sqz \
		postgres

.PHONY: adminer
adminer:
	docker run --rm -ti --network host adminer

# migrate:
# 	migrate -source file://migrations \
# 		-database postgres://postgres:secret@localhost/sqz-data?sslmode=disable up

# migrate-down:
# 	migrate -source file://migrations \
# 		-database postgres://postgres:secret@localhost/sqz-data?sslmode=disable down
