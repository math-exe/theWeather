package util

// DegreesToCardinal converte um angulo em graus para a correspondencia do ponto cardial
// Essa funcao tem que ser exportada ->
func DegressToCardinal(deg int) string {
	dirs := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	// Aqui usamos a logica de conversao
	index := int(float64(deg)/45.0 + 0.5)
	return dirs[index%8]
}
