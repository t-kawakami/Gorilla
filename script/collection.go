package main
import "fmt"

func main() {
	sampleArray()
	sampleSlice()
}

func sampleArray() {
	// 配列初期化
	numbers := [3]int{2,3,4}
	// 通常のfor文で出力
	for index := 0; index < len(numbers); index++ {
		fmt.Println(fmt.Sprintf("value=%d", numbers[index]))
	}
	// 拡張for文
	for _, value := range numbers {
		fmt.Println(fmt.Sprintf("value=%d", value))
	}
	numbers2 := [9]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(numbers2[:])
	fmt.Println(numbers2[1:3])
	fmt.Println(numbers2[:3])
	fmt.Println(numbers2[2:])

	// 配列の長さを指定 → 指定したものがArray
	var values2 [3]int
	// values2 = []int{1,2,3} 長さ指定しないとNG
	values2 = [3]int{1,2,3}
	fmt.Println(values2)
	// values2 = [5]int{1,2,3,4,5} 長さ変えられない
}

func sampleSlice() {
	// 配列初期化
	numbers := []int{2,3,4}
	// 通常のfor文で出力
	for index := 0; index < len(numbers); index++ {
		fmt.Println(fmt.Sprintf("value=%d", numbers[index]))
	}
	// 拡張for文
	for _, value := range numbers {
		fmt.Println(fmt.Sprintf("value=%d", value))
	}

	// 配列の長さを指定せず → 指定しないものがSlice
	var values []int
	fmt.Println(values)
	// values[2] = 100 これはNG
	values = []int{0,2,0,0}
	fmt.Println(values)
	// 長さ変えられる
	values = []int{1,2,3,4,5}
	fmt.Println(values)
}

func sampleMap() {

}