package main
import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonTime struct {
	time.Time
}

func (jsonTime JsonTime) format() string {
	return jsonTime.Time.Format("2006-01-02")
}

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + jsonTime.format() + `"`), nil
}

type Gorilla struct {
	Name string `json:"name"`
	Age int `json:age`
	Children[] Gorilla `json:children`
	BirthDay JsonTime `json:birthday`
}

func main() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	var children []Gorilla
	children = append(children, Gorilla{Name:"Nobita", Age:6, Children:[]Gorilla{}, BirthDay:JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, Gorilla{Name:"Suneo", Age:6, Children:[]Gorilla{}, BirthDay:JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, Gorilla{Name:"Jaian", Age:6, Children:[]Gorilla{}, BirthDay:JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, Gorilla{Name:"Shizuka", Age:6, Children:[]Gorilla{}, BirthDay:JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, Gorilla{Name:"Shusaku", Age:6, Children:[]Gorilla{}, BirthDay:JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})

	gorilla := Gorilla{Name:"Boss", Age:16, Children:children, BirthDay:JsonTime{time.Date(1999,1,18,0,0,0,0, loc)}}
	jsonString, _ := json.MarshalIndent(gorilla, " ", "    ")
	fmt.Printf(string(jsonString))
}