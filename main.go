package main

import (
	"fmt"
	
)
func main(){
   age := 45
   fmt.Println(age <= 50)
   fmt.Println(age >= 50)
   fmt.Println(age == 50)
   fmt.Println(age != 50)

   if age < 30{
      fmt.Println("menor que 30 anos")
}else if age < 40{
   fmt.Println("menor que 40 anos")
}else {
   fmt.Println("Não é menor que 40 anos")
}

names := []string{"Otávio", "Denise", "Fabricio", "Gojo", "Neymar"}

for index, value := range names{
   if index == 0{
      fmt.Println("continue após a posição", index, "e valor", value)
      continue
   }
   if index > 3{
       fmt.Println("sair após", index)
       break
   }
   fmt.Println("valor:", value)
}
 }