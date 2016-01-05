package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "bufio"
)

func main() {
    fileName := "data\\gorilla.txt"
    deleteFile(fileName)
    gorilla(fileName)
    readFile(fileName)
}

func gorilla(fileName string) {
    writeFileAppend("　　　　　　　　　,r\"´⌒｀ﾞ`ヽ", fileName)
    writeFileAppend("　　　　　　　／　,　　　-‐- !､", fileName)
    writeFileAppend("　　　 　　／　{,}f　　-‐- ,,,__､)", fileName)
    writeFileAppend("　　　　／　　 /　　.r'~\"''‐--､)", fileName)
    writeFileAppend("　　,r''\"´⌒ヽ{　　 ヽ　(・)ﾊ(・)}､", fileName)
    writeFileAppend("　/　　　　　　＼　　（⊂｀-'つ）i-､", fileName)
    writeFileAppend("　　　　　　　　　 `}. （__,,ノヽ_ﾉ,ﾉ　 ＼", fileName)
    writeFileAppend("　　　　　　　　　　 l　　 ｀-\" ,ﾉ　　　 ヽ", fileName)
    writeFileAppend("　　　　　　　　　　 }　､､___,j''　　　　　 l\")", fileName)
}

func deleteFile(fileName string) {
    os.Truncate(fileName, 0)
}

func writeFileAppend(text string, fileName string) {
    writeFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, os.ModePerm)

    writer := bufio.NewWriter(writeFile)
    writer.WriteString(text)
    writer.WriteString("\r\n")
    writer.Flush()
}

func readFile(fileName string) {
    contents,err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(contents))

}