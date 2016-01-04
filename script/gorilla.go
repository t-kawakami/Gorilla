// You can edit this code!
// Click here and start typing.
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    gorilla()
    write("gorilla")
    read()
}

func gorilla() {
    fmt.Println("　　　　　　　　　,r\"´⌒｀ﾞ`ヽ")
    fmt.Println("　　　　　　　／　,　　　-‐- !､")
    fmt.Println("　　　 　　／　{,}f　　-‐- ,,,__､)")
    fmt.Println("　　　　／　　 /　　.r'~\"''‐--､)")
    fmt.Println("　　,r''\"´⌒ヽ{　　 ヽ　(・)ﾊ(・)}､")
    fmt.Println("　/　　　　　　＼　　（⊂｀-'つ）i-､")
    fmt.Println("　　　　　　　　　 `}. （__,,ノヽ_ﾉ,ﾉ　 ＼")
    fmt.Println("　　　　　　　　　　 l　　 ｀-\" ,ﾉ　　　 ヽ")
    fmt.Println("　　　　　　　　　　 }　､､___,j''　　　　　 l\")")
}

func write(text string) {
    content := []byte(text)
    ioutil.WriteFile("data\\test.txt", content, os.ModePerm)
}

func read() {
    fmt.Println(ioutil.ReadFile("data\\test.txt"))

}