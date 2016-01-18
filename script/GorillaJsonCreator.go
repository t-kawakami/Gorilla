package main
import (
	"encoding/json"
	"fmt"
	"time"
)

type Gorilla struct {
	Name string `json:"name"`
	Age int `json:age`
	Children[] Gorilla `json:children`
	BirthDay time.Time `json:birthday`
}

func main() {
	gorilla := Gorilla{Name:"Boss", Age:16, Children:[]Gorilla{}, BirthDay:time.Now()}
	jsonString, _ := json.MarshalIndent(gorilla, " ", "    ")
	fmt.Printf(string(jsonString))
}