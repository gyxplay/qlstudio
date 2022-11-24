/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-11-16 09:37:46
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-11-17 10:12:46
 * @FilePath: \Memo_frame\CSV\Csv.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */

package CSV

import (
	"encoding/csv"
	// "fmt"
	"log"
	"os"
)

func WriteCSV(C [][]string) {
	//创建csv文件
	f, err := os.OpenFile("Debt.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	//异步管理
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	w := csv.NewWriter(f)
    w.Write([][]string{{},{}})
	w.Write([][]string{{},{}})
	w.Write([][]string{{},{}})
	//写入数据
	w.WriteAll()
	w.Flush()
}

func ReadCSV() [][]string {
	fileName := "Debt.csv"
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}

	return content
}

func ShowCSV() {
	fileName := "Debt.csv"
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}

}
