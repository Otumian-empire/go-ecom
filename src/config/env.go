package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type environ struct {
	Database           string
	DatabaseUsername   string
	DatabasePassword   string
	DatabaseName       string
	DatabasePort       string
	DatabaseHost       string
	ServerPort         string
	DatabaseDriverName string
	JwtIssuer          string
	JwtSecret          string
	JwtTtl             string
	SslMode            string
}

func GetEnvirons() (environ, error) {
	env, err := godotenv.Read()

	if err != nil {
		log.Println(err)
		return environ{}, fmt.Errorf(ENV_ERROR_MESSAGE)
	}

	return environ{
		Database:           env["DATABASE"],
		DatabaseUsername:   env["DATABASE_USERNAME"],
		DatabasePassword:   env["DATABASE_PASSWORD"],
		DatabaseName:       env["DATABASE_NAME"],
		DatabasePort:       env["DATABASE_PORT"],
		DatabaseHost:       env["DATABASE_HOST"],
		ServerPort:         env["SERVER_PORT"],
		DatabaseDriverName: env["DATABASE_DRIVER_NAME"],
		JwtIssuer:          env["JWT_ISSUER"],
		JwtSecret:          env["JWT_SECRET"],
		JwtTtl:             env["JWT_TTL"],
		SslMode:            env["SSL_MODE"],
	}, nil
}
