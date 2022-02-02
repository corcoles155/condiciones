package main

import "fmt"

var status bool


func main()  {
	status = true

	//Condición if/else
	if status==true {
		fmt.Println(status)
	} else {
		fmt.Println("Es falso")
	}

	//Podemos poder una asignación de valor y una condición en la misma línea
	if status=false; status==true {
		fmt.Println(status)
	} else {
		fmt.Println("Es falso")
	}

	//Condición if/else if
	if status==true {
		fmt.Println(status)
	} else {
		fmt.Println("Es falso")
	}

	//Condición switch, declaramos la variable numero y luego la evaluamos. 
	//No es necesario poner break en case.
	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("El número es mayor que 5")					
	}
}	
