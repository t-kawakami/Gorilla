package main

import (
    "./gorilla"
    "encoding/json"
    "fmt"
    "os"
    "./util"
)

type Goritia struct {
    Id string `json:"id"`
    Turn int `json:"turn"`
    Hand []int `json:"hand"`
}

func main() {
    fmt.Println("########## Gorilla ##########")
    fileName := "data\\gorilla.txt"
    util.DeleteFile(fileName)
    gorilla.DoGorilla(fileName)
    fmt.Println(util.ReadFile(fileName))

    fmt.Println("########## json1 ##########")
    gta := Goritia{Id:"test", Turn:1, Hand:[]int{1, 2, 3, 4, 5}}
    jsonString, _ := json.MarshalIndent(gta, "", "    ")
    fmt.Println(string(jsonString))

    fmt.Println("########## json2 ##########")
    var gt Goritia
    json.Unmarshal([]byte(jsonString), &gt)
    gt.Id = "newTest"
    gt.Turn = 2
    gt.Hand[3] = 6
    jsonString2, _ := json.MarshalIndent(gt, "", "    ")
    fmt.Println(string(jsonString2))

    newFile := "data\\newFile.txt"
    fmt.Println("########## create1 ##########")
    util.WriteFile(string(jsonString2), newFile, os.O_CREATE)
    fmt.Println(util.ReadFile(newFile))
    fmt.Println("########## append1 ##########")
    util.WriteFile(string(jsonString2), newFile, os.O_APPEND)
    fmt.Println(util.ReadFile(newFile))
    fmt.Println("########## append2 ##########")
    util.WriteFile(string(jsonString2), newFile, os.O_APPEND)
    fmt.Println(util.ReadFile(newFile))
    fmt.Println("########## create2 ##########")
    util.WriteFile(string(jsonString2), newFile, os.O_CREATE)
    fmt.Println(util.ReadFile(newFile))
    fmt.Println("########## end ##########")
}

