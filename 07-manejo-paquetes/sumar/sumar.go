package sumar

func Sumar(datos ...float64) (total float64) {
	for _, dato := range datos {
		total += dato
	}
	return
}
