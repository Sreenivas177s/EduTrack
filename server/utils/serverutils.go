package utils

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(filepath.Join("..", ".env")); err != nil {
		panic("Error while loading env file")
	}

}
