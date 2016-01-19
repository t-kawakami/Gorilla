package main


import (
	"encoding/json"
	"fmt"
	"./gorilla"
)

func main() {
	unMarshalTest()
	fmt.Println("#####")
	unMarshalBanana()
	fmt.Println("#####")
	unMarshalGorilla()
	fmt.Println("#####")
	unMarshalGorilla2()
}

// int型の配列は展開できた
func unMarshalTest() {
	var newGoritia gorilla.Goritia
	json.Unmarshal([]byte(`{"id":"test","turn":1,"hand":[1,2,3,4,5]}`), &newGoritia)
	jsonNewStr,_ := json.Marshal(newGoritia)
	fmt.Println(string(jsonNewStr))

	for index, hand := range newGoritia.Hand {
		fmt.Printf("%d:%d\r\n", index, hand)
	}
}

// 別のstructを持つ配列も展開できた
func unMarshalBanana() {
	var cluster gorilla.Cluster
	json.Unmarshal([]byte(`{"id":"a cluster of banana", "banana":[{"sweetLevel":1},{"sweetLevel":2}]}`), &cluster)
	jsonStr,_ := json.Marshal(cluster)
	fmt.Println(string(jsonStr))
}

// 同じ型が入れ子になったJsonをうまくパースできていない
func unMarshalGorilla() {
	var gorilla gorilla.Gorilla
	json.Unmarshal([]byte(`{"name":"Boss","age":16,"children":[{"name":"Nobita","age":6,"children":[],"birthday":"2009-01-01"},{"name":"Suneo","age":6,"children":[],"birthday":"2009-01-01"},{"name":"Jaian","age":6,"children":[],"birthday":"2009-01-01"},{"name":"Shizuka","age":6,"children":[],"birthday":"2009-01-01"},{"name":"Shusaku","age":6,"children":[],"birthday":"2009-01-01"}],"birthday":"1999-01-18"}`), &gorilla)

	parentStr, _ := json.MarshalIndent(gorilla, "", "    ")
	fmt.Println(string(parentStr))

	for _, child := range gorilla.Children {
		childStr,_ := json.MarshalIndent(child, "", "    ")
		fmt.Println(string(childStr))
	}
}

func unMarshalGorilla2() {
	var gorilla gorilla.Gorilla
	json.Unmarshal([]byte(`{"name":"Boss","age":16,"children":[],"birthday":"1999-01-18"}`), &gorilla)

	parentStr, _ := json.MarshalIndent(gorilla, "", "    ")
	fmt.Println(string(parentStr))

	for _, child := range gorilla.Children {
		childStr,_ := json.MarshalIndent(child, "", "    ")
		fmt.Println(string(childStr))
	}
}