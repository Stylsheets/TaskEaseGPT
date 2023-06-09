package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Init() {
	InitViper()
	InitDotenv()
}

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func InitDotenv() {
	err := godotenv.Load()
	if err != nil {
		// if .env file is not found, ignore the error
		if _, ok := err.(*os.PathError); !ok {
			log.Fatal(err)
		}
	}
}

type PineconeCredentials struct {
	URI    string
	ApiKey string
}

func GetPineconeCredentials() PineconeCredentials {
	return PineconeCredentials{
		URI:    os.Getenv("PINECONE_URI"),
		ApiKey: os.Getenv("PINECONE_API_KEY"),
	}
}

type OpenAICredentials struct {
	ApiKey string
}

func GetOpenAICredentials() OpenAICredentials {
	return OpenAICredentials{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	}
}
