package main
import "fmt"

func main() {
	sampleArray()
	sampleSlice()
}

func sampleArray() {
	fmt.Println("######")
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
	fmt.Println("######")
	numbers2 := [9]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(numbers2[:])
	fmt.Println(numbers2[1:3])
	fmt.Println(numbers2[:3])
	fmt.Println(numbers2[2:])

	fmt.Println("######")
	// 配列の長さを指定 → 指定したものがArray
	var values2 [3]int
	// values2 = []int{1,2,3} 長さ指定しないとNG
	values2 = [3]int{1,2,3}
	fmt.Println(values2)
	// values2 = [5]int{1,2,3,4,5} 長さ変えられない

	fmt.Println("######")
	// 配列の代入は値渡し
	values3 := values2
	values2[2] = 5
	fmt.Println(values2)
	fmt.Println(values3)
	values3[2] = 4
	fmt.Println(values2)
	fmt.Println(values3)
}

func sampleSlice() {
	fmt.Println("######")
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

	fmt.Println("######")
	// 配列の長さを指定せず → 指定しないものがSlice
	var values []int
	fmt.Println(values)
	// values[2] = 100 これはNG
	values = []int{0,2,0,0}
	fmt.Println(values)
	// 長さ変えられる
	values = []int{1,2,3,4,5}
	fmt.Println(values)

	fmt.Println("######")
	// Sliceの代入はポインタ渡し
	values2 := values
	values2[3] = 2
	fmt.Println(values)
	fmt.Println(values2)

	fmt.Println("######")
	// arrayからsliceを作る → ポインタ渡し
	array := [3]int{1,2,3}
	slice := array[:]
	slice[1] = 6
	fmt.Println(array)
	fmt.Println(slice)

	fmt.Println("######")
	// sliceからsliceを作る → ポインタ渡し
	slice2 := slice[:]
	slice2[2] = 5
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(slice2)

	// sliceからarrayを作る →このあたりはできなかった
	// var array2 [3]int
	// array2 = slice2[0:3]
	// array2 = slice2[:]
	// array2 = slice2[1:3]

	fmt.Println("######")
	// make関数を使う
	slice3 := make([]int, 3, 3)
	fmt.Println(slice3)
	// append関数 → appendの場合は値渡し
	slice4 := append(slice3, 1, 2, 3, 4, 5)
	fmt.Println(slice3)
	fmt.Println(slice4)
	slice4[0] = 3
	slice4[6] = 1
	fmt.Println(slice3)
	fmt.Println(slice4)

}

func sampleMap() {

}