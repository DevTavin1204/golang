package main
import "fmt"

func main(){
    var ages = [4]int{17, 16, 20, 40}
	nomes := [4]string{"fabiano", "heitor", "maria", "otavio"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] = "Fabricio"
	fmt.Println(nomes)
	// Slice
	var score = []int{100, 200, 300, 400}
	fmt.Println(score)
	score[1] = 2
	fmt.Println(score, len(score), cap(score))
	rangeOne := score[1:3]
	fmt.Println(rangeOne)
	rangeTwo := score[2:]
	fmt.Println(rangeTwo)
	rangeThree := score[:3]
	fmt.Println(rangeThree)

	var superherois = []string{"deadpool", "Homem-aranha", "motoqueiro-fantasma"}
	fmt.Println(superherois)
	superherois = append(superherois, "ben 10", "Gojo")
	fmt.Println(superherois, len(superherois), cap(superherois))
	
}