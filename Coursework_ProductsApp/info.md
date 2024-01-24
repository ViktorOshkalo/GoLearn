commands to run app:

1.
docker-compose up -d

2.
(install migrate)
migrate -path ./migrations -database "mysql://user:password@tcp(localhost:3306)/ProductsAppDb" up

3. install protoc-gen-go (if you need generate from proto)
go install github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

