package provider

import "time"

// WeaterInfo sera a estrutura de dominio
type WeatherInfo struct {
	City        string
	Description string
	Temperature float64 // Celsius
	FetchedAt   time.Time
	TempMin     float64
	TempMax     float64
	WindSpeed   float64
	WindDir     string // Aqui retorna em graus, entao, convertemos em direcao
	Sunrise     time.Time
	Sunset      time.Time
}

// WeatherProvider define o contrato para qualquer provedor de dados de tempo/clima
type WeatherProvider interface {
	FetchByCity(city string) (*WeatherInfo, error)
}
