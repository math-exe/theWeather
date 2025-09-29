package openweather

// apiResponse -> Estrutura que espelha o JSON pela API da OpenWeatherMap
// Mantemos ela separada do modelo de dominio (provider.WeatherInfo)

type apiResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main []struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}
