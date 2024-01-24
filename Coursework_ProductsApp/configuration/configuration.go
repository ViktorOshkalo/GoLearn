package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ConnectionStringTest string = "user:password@tcp(localhost:3306)/ProductsAppDb?parseTime=True"

var SqlProvider string = "mysql"

var User string
var Password string
var ConnectionString string
var GrpcHost string
var GrpcServerPort string

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
	)

	User = os.Getenv("APP_USER")
	Password = os.Getenv("APP_USER_PASSWORD")

	GrpcHost = os.Getenv("GRPC_HOST")
	GrpcServerPort = os.Getenv("GRPC_SERVER_PORT")
}
