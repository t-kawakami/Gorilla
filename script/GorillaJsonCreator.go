package main
import (
	"encoding/json"
	"fmt"
	"time"
	"./gorilla"
	"./jsonTime"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	var children []gorilla.Gorilla
	children = append(children, gorilla.Gorilla{Name:"Nobita", Age:6, Children:[]gorilla.Gorilla{}, BirthDay:jsonTime.JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, gorilla.Gorilla{Name:"Suneo", Age:6, Children:[]gorilla.Gorilla{}, BirthDay:jsonTime.JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, gorilla.Gorilla{Name:"Jaian", Age:6, Children:[]gorilla.Gorilla{}, BirthDay:jsonTime.JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, gorilla.Gorilla{Name:"Shizuka", Age:6, Children:[]gorilla.Gorilla{}, BirthDay:jsonTime.JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})
	children = append(children, gorilla.Gorilla{Name:"Shusaku", Age:6, Children:[]gorilla.Gorilla{}, BirthDay:jsonTime.JsonTime{time.Date(2009,1,1,0,0,0,0,loc)}})

	gorilla := gorilla.Gorilla{Name:"Boss", Age:16, Children:children, BirthDay:jsonTime.JsonTime{time.Date(1999,1,18,0,0,0,0, loc)}}
	jsonString, _ := json.MarshalIndent(gorilla, " ", "    ")
	fmt.Printf(string(jsonString))
}