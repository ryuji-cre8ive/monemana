up:
	docker compose build
	docker compose up

start-local:
	air -c .air.toml

ngrok:
	ngrok http 8080