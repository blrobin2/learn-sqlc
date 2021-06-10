PHONEY: init generate

init:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc init

generate:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate