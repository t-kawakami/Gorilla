package main


import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"./gorilla"
)

func main() {
	goritia := gorilla.Goritia{Id:"test", Turn:1, Hand:[]int{1, 2, 3, 4, 5}}
	jsonStr,_ := json.Marshal(goritia)
	fmt.Println(string(jsonStr))

	var newGoritia gorilla.Goritia
	json.Unmarshal(jsonStr, &newGoritia)
	jsonNewStr,_ := json.Marshal(newGoritia)
	fmt.Println(string(jsonNewStr))

	for hand := range newGoritia.Hand {
		fmt.Println(string(hand))
	}
}

// Gorilla JSONをパースしてChildrenを一行ずつ取り出す
func unMarshalGorilla() {
	jsonString,_ := ioutil.ReadFile("data/gorillaJson.txt")
	var gorilla gorilla.Gorilla
	json.Unmarshal([]byte(jsonString), &gorilla)

	parentStr,_ := json.MarshalIndent(gorilla, " ", "    ")
	fmt.Println(string(parentStr))

	for child := range gorilla.Children {
		childStr,_ := json.MarshalIndent(child, " ", "    ")
		fmt.Println(string(childStr))
	}

}