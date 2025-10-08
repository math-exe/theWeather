package openweather

// apiResponse -> Estrutura que espelha o JSON pela API da OpenWeatherMap
// Mantemos ela separada do modelo de dominio (provider.WeatherInfo)

type apiResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp    float64 `json:"temp"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Sys struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`
}
