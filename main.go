package main

import ("fmt")

var saldo int = 10000
var resposta int

func sacar(valorSaque int) int{
fmt.Println("quanto você deseja sacar?")
    fmt.Scan(&valorSaque)
    if valorSaque > 10000{
      fmt.Println("Saldo insuficiente")
  } 
resposta = saldo - valorSaque
return resposta
}
func depositar(valorDeposito int) int{
  fmt.Println("quanto você deseja depositar?")
    fmt.Scan(&valorDeposito)
  resposta = saldo + valorDeposito
  return resposta
}
func verSaldo(valorConta int){
  fmt.Println("O seu saldo é:", resposta)
}

func main(){
  escolha := "sacar || depositar" 
  fmt.Println("Oque você quer fazer? sacar/depositar")
  fmt.Scan(&escolha)

  if escolha == "sacar"{
    sacar(saldo)

  }else if escolha == "depositar"{
    depositar(saldo)
  }
  verSaldo(resposta)
}