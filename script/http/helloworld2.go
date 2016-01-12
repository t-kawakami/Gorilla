package main
import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Gorilla struct {
	Height int
	Weight int
}

func (gorilla Gorilla) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	text,_ := json.Marshal(gorilla)
	fmt.Fprint(writer, string(text))
}

func main() {
	http.Handle("/", Gorilla{180, 200})
	http.ListenAndServe("localhost:8080", nil)
}


