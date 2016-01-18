package gorilla

import (
	"../jsonTime"
)
type Gorilla struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Children []Gorilla `json:"children"`
	BirthDay jsonTime.JsonTime `json:"birthday"`
}

type Goritia struct {
	Id string `json:"id"`
	Turn int `json:"turn"`
	Hand []int `json:"hand"`
}
