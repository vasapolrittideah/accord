APP_BINARY=api

up:
	docker compose up -d

down:
	docker compose down

up-build: build-api down
	docker compose up --build -d

build-api:
	CGO_ENABLED=0 GOOS=linux go build -o ./dist/apps/${APP_BINARY} ./apps/api/cmd/api/main.go