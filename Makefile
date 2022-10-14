dev:
	@echo "[dev] Running":
	@nodemon --exec go run cmd/api/main.go --signal SIGTERM 