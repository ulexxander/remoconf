.PHONY: all

tools:
	go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	swag init -d ./restapi -g ../cmd/remoconf/main.go -p pascalcase --parseDependency
