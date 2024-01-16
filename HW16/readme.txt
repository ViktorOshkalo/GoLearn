commands to run app:

1.
docker-compose up -d

2.
migrate -path ./migrations -database "mysql://user:password@tcp(localhost:3306)/ProductsAppDb" up

3.
go run main.go