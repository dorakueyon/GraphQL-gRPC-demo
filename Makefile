before: FORCE
	git pull

build: FORCE
	@echo "docker build"
	@docker compose build --no-cache

up:
	@echo "docker up"
	@docker compose up -d --remove-orphans

server: FORCE
	go run ./article/server/server.go




gqlgen: FORCE
	go run github.com/99designs/gqlgen init

client: FORCE
	go run graph/server/server.go

FORCE:
.PHONY: FORCE
