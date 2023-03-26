package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func WriteToDotEnv(openApiKey string) {
	env, _ := godotenv.Unmarshal("OPENAI_API_KEY=" + openApiKey)
	err := godotenv.Write(env, ".env")
	if err != nil {
		log.Println("There was an error writing to the dotenv file")
	}
}
