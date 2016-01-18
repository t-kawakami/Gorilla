package main
import (
	"encoding/json"
	"fmt"
	"time"
)

type jsonTime struct {
	time.Time
}

// formatを指定する。2006-01-02や2006-01-02 03:04:05など。
func (jsonTime jsonTime) format() string {
	return jsonTime.Time.Format("2006-01-02")
}

func (jsonTime jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + jsonTime.format() + `"`), nil
}


type Gorilla struct {
	Name string `json:"name"`
	Age int `json:age`
	Children[] Gorilla `json:children`
	BirthDay jsonTime `json:birthday`
}

func main() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	gorilla := Gorilla{Name:"Boss", Age:16, Children:[]Gorilla{}, BirthDay:jsonTime{time.Date(2015,1,18,0,0,0,0,loc)}}
	jsonString, _ := json.MarshalIndent(gorilla, " ", "    ")
	fmt.Printf(string(jsonString))
}