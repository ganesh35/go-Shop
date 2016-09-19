package lib

import (
	"log"
    "os"
    "fmt"
    "time"
    "encoding/json"
//    "bytes"
)
type GLog struct{
	Destination string
    LogDate string
    LogItems []*LogItem
}

type LogItem struct{
    ItemDate string
    ItemType string
    ItemText string
    FromIp  string
}


func (p *GLog) Info(entry string){
    p.createEntry("Info", entry)
	log.Println("Info||", entry )
}
func (p *GLog) Error(entry string){
    p.createEntry("Error", entry)
	log.Println("Error||", entry)
}
func (p *GLog) Warning(entry string){
    p.createEntry("Warning", entry)
	log.Println("Warning||", entry)
}
func (p *GLog) Critical(entry string){
    p.createEntry("Critical", entry)
    log.Println("Critical||", entry)
}


func (p *GLog) createEntry(entry_type string, entry string){
    t := time.Now()
    timestring := fmt.Sprintf("%d%02d%02d%02d%02d%02d",t.Year(), t.Month(), t.Day(),t.Hour(), t.Minute(), t.Second())

    logItem := new (LogItem)
    logItem.ItemDate = timestring
    logItem.ItemType = entry_type
    logItem.ItemText = entry
    logItem.FromIp = ""

    p.LogItems = append(p.LogItems, logItem)

}

func (p *GLog) Close(logFolder string, logFile string, logFormat string){
    var err error
    t := time.Now()
    timestring := fmt.Sprintf("%d%02d%02d%02d%02d%02d",t.Year(), t.Month(), t.Day(),t.Hour(), t.Minute(), t.Second())
    p.LogDate =timestring
//    log.Println(p.LogItems)


    var fileData []byte
    if logFormat == "json"{
        logFile = logFile + ".json"
        fileData, err = json.MarshalIndent(p.LogItems, "", " ")    
        if err != nil {
            fmt.Println("Log file creation erro : " , err)
            return
        }
    } else {
        logFile = logFile + ".txt"
        fileData = []byte(p.String())
    }
    

    
    f, err := os.OpenFile(logFolder + timestring + "_" +logFile, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        log.Printf("error opening file: %v", err)
    }

    n2, err := f.Write( fileData )
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    defer f.Close()   
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}



func (p *GLog) String() (string){
    output := ""
    for _, item := range p.LogItems {
        stringSlice := item.ItemDate +"||"+ item.ItemType +"||"+  item.ItemText +"||"+ item.FromIp+"\n"
        output = output + stringSlice
    }
    return output;
}

// Usage
/*
var gLog GLog

GLog.Error("This is an error message")

main close
func Close(){
    gLog.Close(LogFolder, LogFile, "json")
}
*/