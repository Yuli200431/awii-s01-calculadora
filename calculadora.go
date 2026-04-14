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