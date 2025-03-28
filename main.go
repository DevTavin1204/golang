package main
import "fmt"

func main(){
	var a int 
	var b int
	fmt.Println("Escolha um numero:")
	fmt.Scan(&a)
	fmt.Println("Escolha outro numero:")
	fmt.Scan(&b)

	fmt.Println("A soma dos numeros é: ", a + b)
	fmt.Println("A subtração dos numeros é: ", a - b)
	fmt.Println("A divisão dos numeros é: ", a / b)
	fmt.Println("A multiplicação dos numeros é: ", a * b)
	fmt.Println("O resto da divisão é:", a % b)
}