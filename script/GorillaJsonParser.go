package main


import (
	"encoding/json"
	"fmt"
	"./gorilla"
	"time"
)

func main() {
	unMarshalTest()
	fmt.Println("#####")
	unMarshalBanana()
	fmt.Println("#####")
	unMarshalGorilla()
	fmt.Println("#####")
	unMarshalGorilla2()
	fmt.Println("#####")
	unMarshalChimpanzee()
	fmt.Println("#####")
	unMarshalChimpanzee2()
	fmt.Println("#####")
	unMarshalTime(`{"name":"Boss","age":16,"birthDay":"2000-01-08T09:00:00Z"}`)
	unMarshalTime(`{"name":"Boss","age":16,"birthDay":"2000-01-08"}`)
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

// 子供がいない場合でも日付をうまくパースできていない
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

// 親と子の型が違う場合はいける？
func unMarshalChimpanzee() {
	type ChimpanzeeChild struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}
	type Chimpanzee struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Children []ChimpanzeeChild `json:"children"`
	}
	var chimpanzee Chimpanzee
	json.Unmarshal([]byte(`{"name":"Boss","age":16,"children":[{"name":"Mary","age":8},{"name":"John","age":4}]}`), &chimpanzee)

	chimpanzeeStr, _ := json.MarshalIndent(chimpanzee, "", "    ")
	fmt.Println(string(chimpanzeeStr))
}

// 親と子の型が同じで日付けを含まない場合
func unMarshalChimpanzee2() {
	type Chimpanzee struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Children []Chimpanzee `json:"children"`
	}
	var chimpanzee Chimpanzee
	json.Unmarshal([]byte(`{"name":"Boss","age":16,"children":[{"name":"Mary","age":8,"children":[]},{"name":"John","age":4,"children":[]}]}`), &chimpanzee)

	chimpanzeeStr, _ := json.MarshalIndent(chimpanzee, "", "    ")
	fmt.Println(string(chimpanzeeStr))
	// うまくいった。Gorillaでうまくいかなかったのは日付けのパースで失敗したことが契機？
}

func unMarshalTime(jsonStr string) {
	type Chimpanzee struct {
		Name string `json:"name"`
		Age int `json:"age"`
		BirthDay time.Time `json:birthDay`
	}

	var chimpanzee Chimpanzee
	json.Unmarshal([]byte(jsonStr), &chimpanzee)

	fmt.Println(chimpanzee.BirthDay.Format("2006-01-02 03:04:05"))
	chimpanzeeStr, _ := json.MarshalIndent(chimpanzee, "", "    ")
	fmt.Println(string(chimpanzeeStr))
}