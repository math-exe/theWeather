package util

// DegreesToCardinal converte um angulo em graus para a correspondencia do ponto cardial
// Essa funcao tem que ser exportada ->
func DegressToCardinal(deg int) string {
	dirs := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	// Aqui usamos a logica de conversao
	index := int(float64(deg)/45.0 + 0.5)
	return dirs[index%8]
}

// GetWeatherEmoji recebe o IconCode da API e retorna o Emoji referente ao clima
// Essa funcao tem que ser exportada ->
func GetWeatherEmoji(iconCode string) string {
	switch iconCode {
	case "01d": // Ceu limpo -> Dia
		return "☀️"
	case "01n": // Ceu limpo -> Noite
		return "🌙"
	case "02d": // Poucas nuvens -> Dia
		return "🌤️"
	case "02n": // Poucas nuvens -> Noite
		return "☁️"
	case "03d", "03n": // Nuvens dispersas -> Dia ou Noite
		return "☁️"
	case "04d", "04n": // Nuvens quebradas -> Dia ou Noite
		return "🌥️"
	case "09d", "09n": // Chuva de banho -> Dia ou Noite
		return "🌧️"
	case "10d": // Chuva -> Dia
		return "🌦️"
	case "10n": // Chuva -> Noite
		return "🌧️"
	case "11d", "11n": // Trovoado -> Dia ou Noite
		return "⛈️"
	case "13d", "13n": // Neve -> Dia ou Noite
		return "❄️"
	case "50d", "50n": // Nevoa -> Dia ou Noite
		return "🌫️"
	default: // Caso o codigo nao tenha sido mapeado, ou nao tenha sido encontrado
		return "❓"
	}
}
