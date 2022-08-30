package initiliazers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(){
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err)
	}
}