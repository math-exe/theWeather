package provider

import "time"

// WeaterInfo sera a estrutura de dominio
type WeatherInfo struct {
	City        string
	Description string
	Temperature float64 // Celsius
	FetchedAt   time.Time
}

// WeatherProvider define o contrato para qualquer provedor de dados de tempo/clima
type WeatherProvider interface {
	FetchByCity(city string) (*WeatherInfo, error)
}
