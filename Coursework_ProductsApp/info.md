commands to run app:

1.
docker-compose up -d

2.
(install migrate)
migrate -path ./migrations -database "mysql://user:password@tcp(localhost:3306)/ProductsAppDb" up

3. generate code from proto
go get github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative catalog.proto



