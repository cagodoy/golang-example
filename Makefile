PORT=5500

dev:
	@echo "[dev] Running":
	@PORT=$(PORT) go run cmd/api/main.go