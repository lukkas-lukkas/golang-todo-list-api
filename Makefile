build:
	docker-compose build

up:
	docker-compose up -d
	docker logs -f golang-todo-list.api