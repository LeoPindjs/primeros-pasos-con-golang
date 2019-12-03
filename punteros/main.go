package main

import (
	"fmt"
)


func duplicar(p *int){
	*p = *p*2
	// *p*=2
}

func main()  {
	
	numero:= 8
	//Copiamos la dirección en memoria de la variable número
	p := &numero
	//Extraemos el valor de p y lo duplicamos
	*p = *p * 2
	fmt.Println(*p)
	// se modifica el valor original de número
	fmt.Println(numero)	
	fmt.Println(*p == numero)	
	
	/*numero := 8
	duplicar(&numero)
	fmt.Println(numero)
	*/

}