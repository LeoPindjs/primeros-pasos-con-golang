package main

import (
	"fmt"

	"github.com/LeoPindjs/primeros-pasos-en-golang/07-manejo-paquetes/saludar"
)

func main() {
	s := saludar.Saludar("Leo")
	fmt.Println(s)
}
