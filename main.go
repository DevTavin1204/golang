package main
import ("fmt")
func main(){
  var numeros = [5]int{}
  var soma int

 fmt.Println("escolha o primeiro numero:")
 fmt.Scan(&numeros[0])
 fmt.Println("escolha o segundo numero:")
 fmt.Scan(&numeros[1])
 fmt.Println("escolha o terceiro numero:")
 fmt.Scan(&numeros[2])
 fmt.Println("escolha o quarto numero:")
 fmt.Scan(&numeros[3])
 fmt.Println("escolha o quinto numero:")
 fmt.Scan(&numeros[4])

soma = numeros[0]+numeros[1]+numeros[2]+numeros[3]+numeros[4]
 fmt.Println("A soma dos numeros é:", soma)
 }