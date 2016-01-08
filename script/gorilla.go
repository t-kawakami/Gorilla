package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "bufio"
)

type Goritia struct {
    Id string `json:"id"`
    Turn int `json:"turn"`
    Hand []int `json:"hand"`
}

func main() {
    fmt.Println("########## Gorilla ##########")
    fileName := "data\\gorilla.txt"
    deleteFile(fileName)
    gorilla(fileName)
    fmt.Println(readFile(fileName))

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
    writeFile(string(jsonString2), newFile, os.O_CREATE)
    fmt.Println(readFile(newFile))
    fmt.Println("########## append1 ##########")
    writeFile(string(jsonString2), newFile, os.O_APPEND)
    fmt.Println(readFile(newFile))
    fmt.Println("########## append2 ##########")
    writeFile(string(jsonString2), newFile, os.O_APPEND)
    fmt.Println(readFile(newFile))
    fmt.Println("########## create2 ##########")
    writeFile(string(jsonString2), newFile, os.O_CREATE)
    fmt.Println(readFile(newFile))
    fmt.Println("########## end ##########")
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

func writeFile(text string, fileName string, flag int) {
    if flag == os.O_APPEND {
        writeFileAppend(text, fileName)
    } else {
        ioutil.WriteFile(fileName, []byte(text), os.ModePerm)
    }
}

func readFile(fileName string) string {
    contents,_ := ioutil.ReadFile(fileName)
    return string(contents)
}

func gorilla(fileName string) {
    writeFileAppend("##################@#####################", fileName)
    writeFileAppend("################@,######################", fileName)
    writeFileAppend("##############@@'#######################", fileName)
    writeFileAppend("############@@##########################", fileName)
    writeFileAppend("###########;@###########################", fileName)
    writeFileAppend("##########@@@;##########################", fileName)
    writeFileAppend("##########.;.+##########################", fileName)
    writeFileAppend("##########  +'##########################", fileName)
    writeFileAppend("#########' ##:##########################", fileName)
    writeFileAppend("#########    :##########################", fileName)
    writeFileAppend("########@  '+,;#########################", fileName)
    writeFileAppend("########`  `@:#@@@######################", fileName)
    writeFileAppend("########`  ''      #####;    @##########", fileName)
    writeFileAppend("########    @        :        ##########", fileName)
    writeFileAppend("########   ,+                 ##########", fileName)
    writeFileAppend("########  `#;       .         ;#########", fileName)
    writeFileAppend("#######. .,#@ ;     . `     ` +#########", fileName)
    writeFileAppend("#######  #+#@@###     ' `#####+#########", fileName)
    writeFileAppend("######@  `+##'.#############@###########", fileName)
    writeFileAppend("######`   ##@  ,.++ '.  ;:``+###########", fileName)
    writeFileAppend("######   @####,. .  ,     ``############", fileName)
    writeFileAppend("######  ,####@ ``     :     ,###########", fileName)
    writeFileAppend("#####@  ,####  .    #'    @.`+##########", fileName)
    writeFileAppend("#####:   #+##`  + @``  @ #    ##########", fileName)
    writeFileAppend("#####.  ;#@#@ , +  @  @ @    `##########", fileName)
    writeFileAppend("######   @##@ ` :   .`,      `@#########", fileName)
    writeFileAppend("######,  @### ` `    `.      `@#########", fileName)
    writeFileAppend("######+   ###+       `.    `` ##########", fileName)
    writeFileAppend("#######   ###;  +    @ ,++'  @##########", fileName)
    writeFileAppend("#######@  ':#    #@#+#@###+;#@##########", fileName)
    writeFileAppend("######## ` ;## ++#########@#############", fileName)
    writeFileAppend("########+   ##.#''#########'############", fileName)
    writeFileAppend("#########   '@. ,.' ,.  :@#@############", fileName)
    writeFileAppend("#########+   ;``., + . '@`:#@###########", fileName)
    writeFileAppend("##########:#:` ,. .:`+..`;#:############", fileName)
    writeFileAppend("############',  +, . , ,.###############", fileName)
    writeFileAppend("############## ' '@: ` #################", fileName)
    writeFileAppend("################+:#,####################", fileName)
    writeFileAppend("#################`, ,'##################", fileName)
    writeFileAppend("###################+`''#################", fileName)
    writeFileAppend("#####################' #################", fileName)
    writeFileAppend("#########,+;',##########################", fileName)
    writeFileAppend("#########:+#',#'### ` ,` :##'########@,#", fileName)
    writeFileAppend("#####+ .''` ' ;@::;@'''''###,',`  `,;'@#", fileName)
    writeFileAppend("#####''' `   ',#,''.    .+##@'''''''''##", fileName)
    writeFileAppend("##### '''''',' #''  ..,,. ,#@',  `''' ##", fileName)
    writeFileAppend("#####@@#''+@''`+'' '''''''+@`##;##@'''##", fileName)
    writeFileAppend("####### ''` ,:' '':''''''':''@# ''  :###", fileName)
    writeFileAppend("####+;''''.####'',++.'' #::'' #,''+ +###", fileName)
    writeFileAppend("####''''''@### '' ##:''+ '''@#@''.#@####", fileName)
    writeFileAppend("###,''''''@##;'' ## '' ###''##.',@######", fileName)
    writeFileAppend("##@@####@#:#,`.@##`'`'#######.`.########", fileName)
    writeFileAppend("##################@#####################", fileName)
}
