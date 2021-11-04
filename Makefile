before: FORCE
	git pull

run-server: FORCE
	go run ./article/server/server.go

FORCE:
.PHONY: FORCE
