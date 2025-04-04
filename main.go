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
}