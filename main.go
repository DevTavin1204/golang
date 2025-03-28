package main
import "fmt"

func main(){
	var usuario = "admin"
	var senha = 1234

	
	var a string
	var b int
	fmt.Println("Escolha um nome de usario:")
	fmt.Scan(&a)
	fmt.Println("Escolha uma senha em numeros:")
	fmt.Scan(&b)

	if a == usuario && b == senha{
		fmt.Println("Você acessou o programa!")
	}else {
		fmt.Println("Acesso negado")
	}
}