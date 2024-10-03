build:
	if [ ! -f .env ]; then cp .env.example .env; fi
	docker-compose build

up:
	make down
	docker-compose up -d
	docker logs -f golang-todo-list.api

down:
	docker-compose down --remove-orphans
