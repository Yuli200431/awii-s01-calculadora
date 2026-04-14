//Paso 1. Imprimir un encabezado decorativo y mostrar la suma de dos números fijos.

package main

import "fmt"
func main() {
	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
	a := 10
	b := 3
	suma := a + b
	fmt.Printf("Suma: %d + %d = %d\n", a, b, suma)
}

//Paso 2. Amplía el programa para que muestre las 4 operaciones básicas (suma, resta, multiplicación, división) usando los mismos valores fijos de 'a' y 'b'.

package main

import "fmt"
func main() {
    fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
	a := 10
	b := 3
	f :=float64(a) / float64(b)
	suma := a + b
	resta := a - b
	multiplicacion := a * b
	fmt.Printf("Suma: %d + %d = %d\n", a, b, suma)
	fmt.Printf("Resta: %d - %d = %d\n", a, b, resta)
	fmt.Printf("Multiplicación: %d * %d = %d\n", a, b, multiplicacion)
	fmt.Printf("División: %d / %d = %.2f\n", a, b, f)

}

//Paso 3. Reemplaza los valores fijos de 'a' y 'b' por entrada del usuario. Permite también que el usuario elija qué operación realizar

package main

import "fmt"

func main() {
    fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
	var num1 float64
	fmt.Print("Digite el primer número: ")
	fmt.Scanln(&num1)
	var num2 float64
	fmt.Print("Digite el segundo número: ")
	fmt.Scanln(&num2)
	var operacion string
	fmt.Print("Seleccione la operación (+, -, *, /): ")
	fmt.Scanln(&operacion)
	
	if operacion == "/" && num2 == 0 {
		fmt.Println("Error: No se puede dividir entre cero.")
		return
	}
	
	var resultado float64
	switch operacion {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		resultado = num1 / num2
	default:
		fmt.Println("Error: Operación no válida.")
		return
	}
	
	fmt.Printf("El resultado de %g %s %g es: %g\n", num1, operacion, num2, resultado)
}


//Paso 4. Agrega dos operaciones científicas adicionales a tu calculadora: potencia y factorial. Ambas deben implementarse manualmente con bucles for, sin usar funciones del paquete math.

package main

import "fmt"

func main() {
	i := 0
	for i < 4 {
		fmt.Println(i)
		i++

    fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
	var num1 float64
	fmt.Print("Digite el primer número: ")
	fmt.Scanln(&num1)
	var num2 float64
	fmt.Print("Digite el segundo número: ")
	fmt.Scanln(&num2)
	var operacion string
	fmt.Print("Seleccione la operación (+, -, *, /,^,! ): ")
	fmt.Scanln(&operacion)

	var resultado float64

	if operacion == "/" && num2 == 0 {
			fmt.Println("Error: No se puede dividir entre cero.")
		return
	}
	
	switch operacion {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		resultado = num1 / num2
	case "^":
		if num2 == 0 {
			resultado = 1
		} else {
			resultado = 1
			for i := 0; i < int(num2); i++ {
				resultado *= num1
			}
		}
	case "!":
		if num1 == 0 {
			resultado = 1
		} else {
			resultado = 1
			for i := 1; i <= int(num1); i++ {
				resultado *= float64(i)
			}
		}					
	default:
		fmt.Println("Error: Operación no válida.")
		return
	}
	
	fmt.Printf("El resultado de %.2f %s %.2f es: %.2f\n", num1, operacion, num2, resultado)
	}

}

//Paso 5. Convierte tu calculadora en una herramienta interactiva que permita realizar varias operaciones seguidas y muestre un historial al final.


package main

import "fmt"

func main() {
	i := 0
	for i < 4 {
		fmt.Println(i)
		i++

    fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
	var num1 float64
	fmt.Print("Digite el primer número: ")
	fmt.Scanln(&num1)
	var num2 float64
	fmt.Print("Digite el segundo número: ")
	fmt.Scanln(&num2)
	var operacion string
	fmt.Print("Seleccione la operación (+, -, *, /,^,! ): ")
	fmt.Scanln(&operacion)

	var resultado float64

	if operacion == "/" && num2 == 0 {
			fmt.Println("Error: No se puede dividir entre cero.")
		return
	}
	
	switch operacion {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		resultado = num1 / num2
	case "^":
		if num2 == 0 {
			resultado = 1
		} else {
			resultado = 1
			for i := 0; i < int(num2); i++ {
				resultado *= num1
			}
		}
	case "!":
		if num1 == 0 {
			resultado = 1
		} else {
			resultado = 1
			for i := 1; i <= int(num1); i++ {
				resultado *= float64(i)
			}
		}					
	default:
		fmt.Println("Error: Operación no válida.")
		return
	}
	
	fmt.Printf("El resultado de %.2f %s %.2f es: %.2f\n", num1, operacion, num2, resultado)
	}


}