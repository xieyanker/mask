package main

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/xieyanker/mask/modules"
	"net/http"
	"io"
	"strings"
	"time"
)

const (
	refreshInterval = 2 * time.Second // 刷新间隔
)

// 初始化屏幕布局
func initScreen() {
	fmt.Print("\033[H\033[2J") // 清屏
	fmt.Printf("\033[%d;0H", 1) // 定位到输出行
}

func main() {
	initScreen()
	ticker := time.NewTicker(refreshInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			showDisplay()
		}
	}
}


func showDisplay() {
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

	var maskIds string
	for _, one := range mask.List {
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

	echo := fmt.Sprintf("Nam\tCur\tChg\tPer\tLow\tHig\tOpe\n")
	for index, _ := range mask.List {
		// Sort by profile order.
		for _, data := range originData.Data {
			if mask.List[index].Id == data.Symbol {
				echo += fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", mask.List[index].Name, data.Current, data.Chg, data.Percent, data.Low, data.High, data.Open)
				break
			}
		}
	}

	var sb strings.Builder
	lines := strings.Split(echo, "\n")
	for index, line := range lines {
		sb.WriteString(fmt.Sprintf("\033[%d;0H\033[K", index + 1))
		sb.WriteString(line)
	}
	fmt.Print(sb.String())
}

