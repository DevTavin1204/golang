package main
import ("fmt"
         "strings"
         "sort"
		)
func main(){
   greeting := "Hello my friends"
   fmt.Println(strings.Contains(greeting, "friends"))
   fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
   fmt.Println(strings.ToUpper(greeting))
   fmt.Println(strings.Index(greeting, "my"))
   fmt.Println(strings.Split(greeting, "friends"))

   ages := []int {50, 80, 10}
   sort.Ints(ages)
   fmt.Println(ages)
   index := sort.SearchInts(ages, 80)
   fmt.Println(index)
   names := []string{"Otávio", "Denise", "Fabricio"}
   sort.Strings(names)
   fmt.Println(names)
   fmt.Println(sort.SearchStrings(names, "Otávio"))
   x :=0

   for x < 5 {
      fmt.Println(x)
      x++
   }

   for i:=0; i <5; i++{
      fmt.Println("For 2: ", i)
   }
   for i:=0; i <len(names); i++{
      fmt.Println(names[i])
   }
   for index, value := range names{
      fmt.Println("O indice é:", index, "e o valor é", value)
   }
   for index, value := range ages{
      fmt.Println("o indice é:", index, " e o valor é:", value)
   }


   superherois := []string{"Deadpool", "Homem aranha", "batman"}

   for index, value := range superherois{
      fmt.Println("O numero do heroi: ", index, "O nome do heroi: ", value)
   }
}