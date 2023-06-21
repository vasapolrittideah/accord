APP_BINARY=server

up:
	docker compose up -d

down:
	docker compose down

up-build: build-server down
	docker compose up --build -d

build-server:
	CGO_ENABLED=0 GOOS=linux go build -o tmp/${APP_BINARY} ./cmd/api/main.go