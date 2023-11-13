package envreader

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("")
	}
    database_url := os.Getenv("DATABASE_URL")

    fmt.Println(database_url)

}