# docker
create-network:
	docker network create go_api_network
docker-up:
	docker-compose -p go-starter-kit-api up -d
docker-down:
	docker-compose -p go-starter-kit-api down

# migration
goose-create:
	goose -dir src/migrations/postgres -s create $(name) sql