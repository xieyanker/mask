package main

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/xieyanker/mask/modules"
	"net/http"
	"io"
	"strings"
)

func main() {
	jsonStr, err := os.ReadFile("conf/mask.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var mask modules.Mask
	err = json.Unmarshal(jsonStr, &mask)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	maskMap := make(map[string]string)
	var maskIds string
	//fmt.Printf("%+v\n", mask)
	for _, one := range mask.List {
		maskMap[one.Id] = one.Name
		maskIds = maskIds + one.Id + ","
	}
	maskIds = strings.TrimSuffix(maskIds, ",")

	// call url
	response, err := http.Get("https://stock.xueqiu.com/v5/stock/realtime/quotec.json?symbol=" + maskIds)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var originData modules.OriginData
	err = json.Unmarshal(body, &originData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if originData.ErrorCode != 0 {
		fmt.Println(originData.ErrorDescription)
		os.Exit(1)
	}

	fmt.Printf("Nam\tCur\tChg\tPer\tLow\tHig\tOpe\n")
	for _, data := range originData.Data {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", maskMap[data.Symbol], data.Current, data.Chg, data.Percent, data.Low, data.High, data.Open)
	}
}
