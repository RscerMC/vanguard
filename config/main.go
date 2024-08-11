package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Token string

	MongoURI string

	Host     string
	Port     int
	Password string
	DB       int

	DeveloperID string

	NormalColor int
	ErrorColor  int
)

func init() {
	godotenv.Load()

	Token = os.Getenv("VANGUARD_TOKEN")
	MongoURI = os.Getenv("MONGO_URI")
	Host = os.Getenv("REDIS_HOST")
	Port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	Password = os.Getenv("REDIS_PASSWORD")
	DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	DeveloperID = os.Getenv("DEVELOPER_ID")
	NormalColor, _ = strconv.Atoi(os.Getenv("NORMAL_COLOR"))
	ErrorColor, _ = strconv.Atoi(os.Getenv("ERROR_COLOR"))
}
