package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DiscordToken string = "DISCORD_TOKEN"
	BotPrefix    string = "BOT_PREFIX"
	LogLevel     string = "LOG_LEVEL"
)

func init() {
	err := godotenv.Load("C:/Users/psuri/Dev/Golang/DiscordBotRe/discord-bot/configs/.env")
	if err != nil {
		log.Fatal("Failed to load .env file: " + err.Error())
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not found", key)
	}
	return value
}
