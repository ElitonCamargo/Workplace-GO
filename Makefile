# Nome do container do MySQL no Docker
CONTAINER_NAME=WP_mysql
DB_USER=root
DB_PASS=root
DB_NAME=sistema_WP

# Caminho para os arquivos SQL
SCHEMA=./MySQL/schema.sql
SEED=./MySQL/seed.sql

# Apaga e recria apenas o schema
reset-schema:
	docker exec -i $(CONTAINER_NAME) mysql -u $(DB_USER) -p$(DB_PASS) < $(SCHEMA)

# Apaga e recria o schema e popula com dados iniciais
reset-db:
	docker exec -i $(CONTAINER_NAME) mysql -u $(DB_USER) -p$(DB_PASS) < $(SCHEMA)
	docker exec -i $(CONTAINER_NAME) mysql -u $(DB_USER) -p$(DB_PASS) $(DB_NAME) < $(SEED)

# Executa apenas o seed (dados iniciais)
seed:
	docker exec -i $(CONTAINER_NAME) mysql -u $(DB_USER) -p$(DB_PASS) $(DB_NAME) < $(SEED)

# Comandos para iniciar os containers Docker
start:
	docker compose up -d

# Comando para construir e iniciar os containers Docker
build:
	docker compose build --no-cache

# Comando para parar e remover os containers Docker
stop:
	docker compose down

# Comando para visualizar os logs do container da API
logs:
	docker compose logs -f api

# Comando para verificar o status dos containers Docker
status:
	docker ps -a