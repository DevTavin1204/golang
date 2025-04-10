package main

import (
	"fmt"
	
)
func main(){
     var saldo int = 20000
     var inserir int 
     var soma int
     var saque int
     var deposito int
     var resposta string = "saque || deposito"
    
     fmt.Println("qual valor você gostaria de inserir?")
     fmt.Scan(&inserir)
     soma = saldo + inserir
     fmt.Println("o seu saldo atual é:", soma)
     fmt.Println("Oque você deseja fazer? saque/deposito")
     fmt.Scan(&resposta)
  
    var operacaoSa int
    var operacaoDe int
     
     if resposta == "saque"{
         fmt.Println("Quanto você deseja sacar?")
         fmt.Scan(&saque)
         if saque > 20000{
            fmt.Println("Saldo insuficiente")
         }
         operacaoSa = soma - saque
         fmt.Println("O saldo da conta é:", operacaoSa)
     }else if resposta == "deposito"{
      fmt.Println("Quanto você deseja depositar?")
      fmt.Scan(&deposito)
      operacaoDe = soma + deposito
      fmt.Println("O saldo da conta é:", operacaoDe)
     }else {
      fmt.Println("Você saiu da conta")
     }
   }
