before: FORCE
	git pull

server: FORCE
	go run ./article/server/server.go

gqlgen: FORCE
	go run github.com/99designs/gqlgen init

client: FORCE
	go run graph/server/server.go

FORCE:
.PHONY: FORCE
