package main

import ("fmt")

func main(){
  estoque := make(map[string]int)
   
   estoque["Coxinha"] = 15
   estoque["Pão de queijo"] = 10
   estoque["Refrigeirante"] = 20

  for produtos, quantidade := range estoque {
    fmt.Println("Produto, Quantidade:", produtos, quantidade) 
   }
   
  estoque["Coxinha"] = 13
  estoque["Pão de queijo"] = 9
   
   for produtos, quantidade := range estoque {
    fmt.Println("Produto, Quantidade:", produtos, quantidade) 
    }
 }