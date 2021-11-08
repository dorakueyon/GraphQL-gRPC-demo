before: FORCE
	git pull

build: FORCE
	@echo "docker build"
	@docker compose build --no-cache

up:
	@echo "docker up"
	@docker compose up -d --remove-orphans

server: FORCE
	go run ./article/main.go

health_check: FORCE
	go run ./article/health_chec/health_check.go

protoc: FORCE
	protoc ./article/article.proto --go_out=plugins=grpc:.

gqlgen_init: FORCE
	go run github.com/99designs/gqlgen init

gqlgen: FORCE
	#rm ./graph/resolver.go
	#rm ./graph/schema.resolvers.go
	go run github.com/99designs/gqlgen

FORCE:
.PHONY: FORCE
