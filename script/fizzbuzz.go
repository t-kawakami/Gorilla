package main
import "fmt"

func main() {
	var array [100]int
	for i := range array {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(fmt.Sprintf("%d:fizzbuzz", i))
		} else if i % 3 == 0 {
			fmt.Println(fmt.Sprintf("%d:fizz", i))
		} else if i % 5 == 0 {
			fmt.Println(fmt.Sprintf("%d:buzz", i))
		}
	}
}
