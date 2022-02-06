package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xuri/excelize/v2"
)

var titles = []string{}

func main() {

	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, _ := f.GetRows("Sheet1")
	cols, _ := f.GetCols("Sheet1")
	//ff := len(rows)

	for i := 0; i < len(cols); i++ {
		titles = append(titles, cols[i][0])
	}
	maps := make([]map[string]string, len(rows))

	//lastarray := make([][]string, ee, ff)
	for j := 0; j < len(rows); j++ {
		new_array := []string{}
		//fmt.Println("-----------")
		for title := range titles {
			//line := fmt.Sprintf("%s:%s", titles[title], rows[j][title])
			//d.title = titles[title]
			//d.datas = rows[j][title]
			new_array = append(new_array, titles[title])
			new_array = append(new_array, rows[j][title])
		}

		elementMap := make(map[string]string)
		for i := 0; i < len(new_array); i += 2 {
			elementMap[new_array[i]] = new_array[i+1]

		}
		maps[j] = elementMap

	}
	json_content, _ := json.MarshalIndent(maps, "", "  ")
	//fmt.Println(string(json_content))
	_ = ioutil.WriteFile("export.json", json_content, 0644)

}
