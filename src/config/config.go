package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	UrlDatabase = ""
	ApiPort     = 0
	SecretKey   []byte
)

func LoadEnv() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		ApiPort = 5500
	}

	UrlDatabase = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		url.QueryEscape(os.Getenv("DB_USER")),
		url.QueryEscape(os.Getenv("DB_PASSWORD")),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
