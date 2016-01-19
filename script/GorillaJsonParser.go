package main


import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"./gorilla"
)

func main() {
	unMarshalTest()
	fmt.Println("#####")
	unMarshalGorilla()
}

func unMarshalTest() {
	var newGoritia gorilla.Goritia
	json.Unmarshal([]byte(`{"id":"test","turn":1,"hand":[1,2,3,4,5]}`), &newGoritia)
	jsonNewStr,_ := json.Marshal(newGoritia)
	fmt.Println(string(jsonNewStr))

	for index, hand := range newGoritia.Hand {
		fmt.Printf("%d:%d\r\n", index, hand)
	}
}

// Gorilla JSONをパースしてChildrenを一行ずつ取り出す
func unMarshalGorilla() {
	jsonString,_ := ioutil.ReadFile("data/gorillaJson.txt")
	var gorilla gorilla.Gorilla
	json.Unmarshal([]byte(jsonString), &gorilla)

	parentStr,_ := json.MarshalIndent(gorilla, " ", "    ")
	fmt.Println(string(parentStr))

	for _, child := range gorilla.Children {
		childStr,_ := json.MarshalIndent(child, " ", "    ")
		fmt.Println(string(childStr))
	}

}