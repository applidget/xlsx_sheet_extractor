# Excel sheet extractor

This simple tool written in Go uses the [xls](https://xgithub.com/tealeg/xlsx) to :
1. Open a spreadsheet (xlsx)
2. List all sheets, and display the number of rows in each of them
3. Print the results as YML to stdout

## Installation

`go get github.com/tealeg/xlsx`
`go build`

Then run it as 

```
xlsx_sheet_extractor some_spreadsheet.xlsx
```
