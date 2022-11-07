
lint:
	go install golang.org/x/lint/golint@latest
	@make run-on-our-code-directories ARGS="golint"

build:generate-swagger
	go build -o resolver && ./resolver

run:build
	./featws-resolver-bridge

generate-swagger:
#   Install swag on https://github.com/swaggo/swag
	swag i

deps:
	go mod tidy