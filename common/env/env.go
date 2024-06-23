package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	API_PORT string
)

func Init() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if API_PORT = os.Getenv("API_PORT"); API_PORT == "" {
		panic(1)
	}
}
