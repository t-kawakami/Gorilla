package main
import (
	"net/http"
	"text/template"
)

type Gorilla struct {
	Height int
	Weight int
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	gorilla := Gorilla{175, 200}
	template1,_ := template.New("new").Parse("Height:{{.Height}},Weight:{{.Weight}}")

	template1.Execute(writer, gorilla)
}