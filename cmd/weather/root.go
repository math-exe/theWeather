package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/math-exe/theWeather/internal/provider"
	"github.com/math-exe/theWeather/internal/provider/openweather"
	"github.com/math-exe/theWeather/internal/service"
	"github.com/math-exe/theWeather/internal/util"
	"github.com/math-exe/theWeather/pkg/config"
	"github.com/spf13/cobra"
)

// newRootCmd cria o comando raiz da aplicacao
// Nessa parte que a "injecao de dependencias" acontece
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "theWeather [cidade]",
		Short: "Busca o tempo atual para uma cidade.",
		Long:  `Uma CLI Simples e rapida para consultar o tempo/clima em qualquer cidade do mundo.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// 1. Primeiro carregamos as configuracoes
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("erro ao carregar configuracao: %w", err)
			}
			if cfg.APIKey == "" {
				return fmt.Errorf("apiKey nula ou nao encontrada. Defina a variavel de ambiente OPENWEATHER_API_KEY no .env")
			}

			// 2. Agora montamos as dependencias (Provider -> Service)
			weatherProvider := openweather.NewClient(cfg.APIKey)
			weatherSvc := service.NewWeatherService(weatherProvider)

			// 3. Executamos a logica do projeto
			city := strings.Join(args, " ")
			weather, err := weatherSvc.GetWeatherByCity(city)
			if err != nil {
				return fmt.Errorf("nao foi possivel obter o tempo para '%s':'%w'", city, err)
			}

			// 4. E por fim, demonstramos os dados
			displayWeather(weather)

			return nil
		},
	}
	return cmd
}

// displayWeather mostra o resultado no terminal(CLI)
func displayWeather(w *provider.WeatherInfo) {
	cityColor := color.New(color.FgHiYellow, color.Bold)
	labelColor := color.New(color.FgWhite)
	valueColor := color.New(color.FgHiCyan)

	emoji := util.GetWeatherEmoji(w.IconCode)

	fullDescription := fmt.Sprintf("%s %s", emoji, w.Description)

	cityColor.Printf("Clima em %s\n", w.City)
	fmt.Println(strings.Repeat("-", 38))

	// Usando padding para alinhar os valores
	labelColor.Printf("%-22s", "Descricao:")
	valueColor.Printf("%s\n", fullDescription)

	labelColor.Printf("%-22s", "Temperatura atual:")
	valueColor.Printf("%.1f°C\n", w.Temperature)

	labelColor.Printf("%-22s", "Temperatura minima:")
	valueColor.Printf("%.1f°C\n", w.TempMin)

	labelColor.Printf("%-22s", "Temperatura maxima:")
	valueColor.Printf("%.1f°C\n", w.TempMax)

	labelColor.Printf("%-22s", "Velocidade do vento:")
	valueColor.Printf("%.2f km/h\n", w.WindSpeed)

	labelColor.Printf("%-22s", "Direcao do vento:")
	valueColor.Printf("%s\n", w.WindDir)

	fmt.Println(strings.Repeat("-", 38))
}
