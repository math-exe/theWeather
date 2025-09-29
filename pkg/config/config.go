package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config armazena todas as configuracoes da aplicacao
type Config struct {
	APIKey string
}

// Load carrega as configuracoes das variaveis de ambiente
// Em ambiente DEV, ele primeiro carrega o arquivo .env
func Load() (*Config, error) {
	// godotenv.Load() vai carregar o .env, mas cuidado que ele nao gera erro caso o arquivo nao exista
	// Para producao isso esta otimo, mas quando retornar ao ambiente de teste, adicionar o log novamente
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	return &Config{
		APIKey: apiKey,
	}, nil
}
