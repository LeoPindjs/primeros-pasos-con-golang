package main

import "fmt"

type Docente struct {
	nombre string
	materia string
	jornada string
}

func print(i uint,n,m,j string){
	
	s1:= "==== Docente %d ====\n"	
	s2:= "Nombre: %s \nMateria: %s\nJornada: %s\n"
	s3:= "=================="
	
	fmt.Printf(s1,i)
	fmt.Printf(s2,n,m,j)
	fmt.Println(s3)
}

func getDocentes() (Docentes map[uint]Docente) {

	Docentes = map[uint]Docente{
		1:Docente{
			"Pedro",
			"Ingles",
			"Mañana",
		},
		2:Docente{
			"Carlos",
			"Informática",
			"Mañana",
		},
		3:Docente{
			"Andres",
			"Física",
			"Tarde",
		},
		4: Docente{
			"Andres",
			"Física",
			"Tarde",
		},
	}

	return
	
	/*
	Docentes = make(map[uint]Docente)
	d1:= Docente{
		"Pedro",
		"Ingles",
		"Mañana",
	}
	
	d2:= Docente{
		"Carlos",
		"Informática",
		"Mañana",
	}
	
	d3:= Docente{
		"Fabian",
		"Cálculo",
		"Tarde",
	}
	
	d4:= Docente{
		"Andres",
		"Física",
		"Tarde",
	}
	
	Docentes[1] = d1
	Docentes[2] = d2
	Docentes[3] = d3
	Docentes[4] = d4

	return Docentes
	*/
}

func main(){
	
	docentes := getDocentes()
	
	for i,docente := range docentes {
		n:= docente.nombre
		m:= docente.materia
		j:= docente.jornada
		print(i,n,m,j)	
	}		
}