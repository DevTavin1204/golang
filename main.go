package main

import ("fmt")
 var resposta string
 var media int

 func analisarNotas(nota1, nota2 float64)(float64, string){
  media := (nota1 + nota2) / 2
  if media >= 6{
   return media, "Aprovado"
  }else {
   return media, "Reprovado"
  }
}
func main(){
   media, resposta := analisarNotas(7.5, 5.5)
   fmt.Println("Media:", media)
   fmt.Println("Resultado", resposta)
 }