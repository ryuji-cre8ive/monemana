up:
	docker compose build
	docker compose up

start-local:
	air -c .air.toml

ngrok:
	ngrok http 8080

migration-up:
	sql-migrate up
# HINT: if you want to make migration file, you can use this command
# sql-migrate new <migration_name>
