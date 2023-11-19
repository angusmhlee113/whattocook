package envreader

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot load .env file")
	}
    database_url := os.Getenv("DATABASE_URL")
    fmt.Println(database_url)
	return database_url
}