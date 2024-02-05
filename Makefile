.PHONY: run
run:
	go run ./cmd/server/main.go

.DEFAULT_GOAL := run