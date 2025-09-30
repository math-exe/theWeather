package service

import "github.com/math-exe/theWeather/internal/provider"

// WeatherService contem a logica de negocios relacionada ao tempo
// Aqui se torna uma dependencia da interface WeatherProvider, nao de uma implementacao concreta
type WeatherService struct {
	provider provider.WeatherProvider
}

// NewWeatherService cria uma nova instancia de servico do tempo/clima
func NewWeatherService(p provider.WeatherProvider) *WeatherService {
	return &WeatherService{
		provider: p,
	}
}

// GetWeatherCity busca os dados de tempo atraves do provedor
// -> Pensar em adicionar uma logica de cache ou combinar varios dados de outras fontes
func (s *WeatherService) GetWeatherByCity(city string) (*provider.WeatherInfo, error) {
	return s.provider.FetchByCity(city)
}
