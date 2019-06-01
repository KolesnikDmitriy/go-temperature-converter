package convert

// FtoC convert fahrenheit to celsius
func FtoC(t float64) float64 {
	return (t - 32) * 5 / 9
}

// CtoF convert celsius to fahrenheit
func CtoF(t float64) float64 {
	return t * 9 / 5 + 32
}

// KtoC convert celsius to kelvin
func KtoC(t float64) float64 {
	return t - 273.15
}

// CtoK convert celsius to kelvin
func CtoK(t float64) float64 {
	return t + 273.15
}

// KtoF convert celsius to kelvin
func KtoF(t float64) float64 {
	return (t - 273.15) * 9 / 5 + 32
}

// FtoK convert celsius to kelvin
func FtoK(t float64) float64 {
	return (t - 32) * 5 / 9 + 273.15
}