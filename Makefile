
# Comandos úteis para desenvolvimento e execução

build:
	docker compose build

up:
	docker compose up

down:
	docker compose down

restart:
	docker compose down -v && docker compose up --build

ps:
	docker compose ps

logs:
	docker compose logs -f

migrate:
	docker compose exec db psql -U postgres -d deputadosdb -f /docker-entrypoint-initdb.d/init.sql
