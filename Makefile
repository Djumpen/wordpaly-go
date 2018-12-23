start:
	docker-compose up

gendoc:
	swagger generate spec -b ./cmd -o ./swaggerui/swagger.json