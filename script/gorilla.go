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
    fileName := "data\\gorilla.txt"
    deleteFile(fileName)
    gorilla(fileName)
    readFile(fileName)

    gta := Goritia{Id:"test", Turn:1, Hand:[]int{1, 2, 3, 4, 5}}
    jsonString, _ := json.MarshalIndent(gta, "", "    ")
    fmt.Println(string(jsonString))

    var gt Goritia
    json.Unmarshal([]byte(jsonString), &gt)
    gt.Id = "newTest"
    gt.Turn = 2
    gt.Hand[3] = 6
    jsonString2, _ := json.MarshalIndent(gt, "", "    ")
    fmt.Println(string(jsonString2))
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
