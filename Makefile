start:
	docker-compose up mysqldb app

migrate:
	docker-compose up --abort-on-container-exit migration mysqldb

gendoc:
	swagger generate spec -b ./cmd/app -i ./doc/swagger_misc.yaml -o ./doc/swaggerui/swagger.json --scan-models

gendoc_yaml:
	swagger generate spec -b ./cmd/app -i ./doc/swagger_misc.yaml -o ./doc/swagger.yaml --scan-models

gendocall: gendoc gendoc_yaml
