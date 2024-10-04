build:
	if [ ! -f .env ]; then cp .env.example .env; fi
	docker-compose build

up:
	docker-compose up --force-recreate

down:
	docker-compose down --remove-orphans

stack-up:
	docker-compose -f docker-compose-stack.yaml up -d --force-recreate

stack-down:
	docker-compose -f docker-compose-stack.yaml down
