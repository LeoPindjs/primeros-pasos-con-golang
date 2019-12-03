package main

import (
	"fmt"
	"math"
)

type Figuras interface{
	area() float64
	perimetro() float64
}

type Circulo struct{
	radio float64
}

type Rectangulo struct {
	largo, ancho float64
}

func (c Circulo) area() float64{
	return  (math.Pi) * (c.radio*c.radio) 
}

func ( r Rectangulo) area() float64{
	return r.largo*r.ancho
}

func (c Circulo) perimetro() float64{
	return 2 * (math.Pi) * c.radio 
}

func ( r Rectangulo) perimetro() float64{
	return 2*r.largo + 2*r.ancho
}

func Calculos(figura Figuras) (a , p float64){
	a,p = figura.area(),figura.perimetro()
	return
}

func print(a,p float64) {
	fmt.Println("=============================")
	fmt.Printf("Area : %f\nPerímetro: %f\n",a,p)
	fmt.Println("=============================")
}
	
func main() {
	
	c1:= Circulo{4.5}
	r1:= Rectangulo{4.5,4.8}	
	figuras:= []Figuras{c1,r1}
	
	for _,figura := range figuras{
		a,p := Calculos(figura)
		print(a,p)
		
	}

}