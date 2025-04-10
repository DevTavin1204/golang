package main

import (
	"fmt"
	
)
func main(){
     var age int
     fmt.Println("qual é sua idade?")
     fmt.Scan(&age)

     if age < 18 {
      fmt.Println("Você é menor de idade")
     }else if age <= 60{
      fmt.Println("Você é adulto")
     }else {
      fmt.Println("Você é idoso")
     }
 }