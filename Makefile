start:
	docker-compose up

gendoc:
	swagger generate spec -b ./cmd -i ./doc/swagger_misc.yaml -o ./doc/swaggerui/swagger.json --scan-models

gendoc_yaml:
	swagger generate spec -b ./cmd -i ./doc/swagger_misc.yaml -o ./doc/swagger.yaml --scan-models

gendocall: gendoc gendoc_yaml
