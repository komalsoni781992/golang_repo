package functionsintro

// Div - divides numerator to denominator
func Div(numerator int64, denominator int64) int64 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
