package env

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func GetNewsTokenAPI() string {
	return os.Getenv("NEWS_API_KEY")
}

func GetNewsURL() string {
	return os.Getenv("NEWS_API_URL")
}