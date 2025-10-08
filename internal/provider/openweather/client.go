package openweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/math-exe/theWeather/internal/provider"
	"github.com/math-exe/theWeather/internal/util"
)

const apiURL = "http://api.openweathermap.org/data/2.5/weather"

// Aqui o Client vai implementar a interface provider.WeatherProvider
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient cria uma nova instancia do cliente da OpenWeatherMap
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// FetchByCity busca os dados do tempo para uma cidade
func (c *Client) FetchByCity(city string) (*provider.WeatherInfo, error) {
	escapeCity := url.QueryEscape(city)

	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric&lang=pt_br", apiURL, escapeCity, c.apiKey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("falha ao realizar a requisicao: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou um erro: %s", resp.Status)
	}

	var data apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("falha ao decodificar a resposta json: %w", err)
	}

	if len(data.Weather) == 0 {
		return nil, fmt.Errorf("resposta da api nao contem dados de tempo")
	}

	// Mapeia a resposta da API para o nosso modelo de dominio interno
	// Essa etapa aqui se torna crucial para o desacoplamento
	weatherInfo := &provider.WeatherInfo{
		City:        data.Name,
		Description: strings.Title(data.Weather[0].Description),
		IconCode:    data.Weather[0].Icon,
		Temperature: data.Main.Temp,
		TempMin:     data.Main.TempMin,
		TempMax:     data.Main.TempMax,
		WindSpeed:   data.Wind.Speed * 3.6,
		WindDir:     util.DegressToCardinal(data.Wind.Deg),
		Sunrise:     time.Unix(data.Sys.Sunrise, 0),
		Sunset:      time.Unix(data.Sys.Sunset, 0),
		FetchedAt:   time.Now(),
	}

	return weatherInfo, nil
}
