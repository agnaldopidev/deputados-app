.PHONY: build up up-detach start stop down down-volumes logs test clean up-pull help

build:
	docker-compose build

up: build
	docker-compose up

up-detach: build
	docker-compose up -d

start:
	docker-compose start

stop:
	docker-compose stop

down:
	docker-compose down

down-volumes:
	docker-compose down -v

logs:
	docker-compose logs -f

test:
	go test ./...

clean:
	docker image prune -f

# Nova target para pull + up com remoção de órfãos
up-pull:
	docker-compose up --pull --remove-orphans

help:
	@echo "Comandos disponíveis:"
	@echo "  make build          # Build das imagens Docker"
	@echo "  make up             # Build + sobe containers (foreground)"
	@echo "  make up-detach      # Build + sobe containers (background)"
	@echo "  make start          # Inicia containers já criados"
	@echo "  make stop           # Para containers"
	@echo "  make down           # Para e remove containers e redes"
	@echo "  make down-volumes   # Para e remove containers, redes e volumes"
	@echo "  make logs           # Logs em tempo real"
	@echo "  make test           # Executa testes Go"
	@echo "  make clean          # Remove imagens dangling"
	@echo "  make up-pull        # Puxa imagens novas + sobe containers removendo órfãos"
	@echo "  make help           # Mostra essa ajuda"
