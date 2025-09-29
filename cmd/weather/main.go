package main

// Structs para mapear a resposta JSON da API do OpenWheaterMap
type WeatherResponse struct {
	Name    string        `json:"name"`
	Main    MainData      `json:"main"`
	Wheater []WheaterData `json:"wheater"`
	Wind    WindData      `json:"wind"`
}

type MainData struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Humidity  float64 `json:"humidity"`
}

type WheaterData struct {
	Description string `json:"description"`
}

type WindData struct {
	Speed string `json:"speed"`
}
