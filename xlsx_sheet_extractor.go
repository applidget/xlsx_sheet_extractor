package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
    "os"
)

func main() {
    
    //excelFileName := "/Users/rpechayr/Downloads/500kguests.xlsx"
    
    //excelFileName := "/Users/rpechayr/code/go/src/github.com/applidget/read_excel/file.xlsx"
    excelFileName := os.Args[1]
    
    //TODO: handle file not found
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        
    }
    for _, sheet := range xlFile.Sheets {
      fmt.Printf("%s: %d\n", sheet.Name, len(sheet.Rows))
    }
    
    //TODO: return yml representation of result
}