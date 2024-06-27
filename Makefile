up:
	docker compose build
	docker compose up

start-local:
	go run main.go

ngrok:
	ngrok http 8080

db-dev-up:
	sql-migrate up -env="dev"
# HINT: if you want to make migration file, you can use this command
# sql-migrate new <migration_name>

db-prod-up:
	sql-migrate up -env="prod"

db-dev-down:
	sql-migrate down -env="dev"

db-prod-down:
	sql-migrate down -env="prod"

