package main

import (
	"fmt"
	"os"
    "crypto/sha256"
    "encoding/hex"
)

type Memory struct{
	date string
	Context string
}

func main(){
	diary := make(map[int]Memory)

	diary[1] = Memory{date:"12/19",Context:"喜欢你"}
	diary[2] = Memory{date:"12/20",Context:"依旧喜欢你"}
	diary[3] = Memory{date:"12/21",Context:"还是很喜欢你"}

	fmt.Println("今天的状态是:")
	for day,m := range diary {

		record := fmt.Sprintf("%d%s%s",day,m.date,m.Context)
		h := sha256.New()
		h.Write([]byte(record))

		hash := h.Sum(nil)
		hashString := hex.EncodeToString(hash)

		fmt.Printf("\033[35m第 %d 天\033[0m | 内容: %s | \033[32m指纹: %s\033[0m\n", 
                   day, m.Context, hashString)
	}
	 
	aa,err := os.Create("love.txt")
	if err != nil{
		fmt.Println("Error..",err)
		return
	}
	defer aa.Close()

	for day,m := range diary {

		record := fmt.Sprintf("%d%s%s",day,m.date,m.Context)
		g := sha256.New()
		g.Write([]byte(record))

		hash := g.Sum(nil)
		hashString := hex.EncodeToString(hash)

		line := fmt.Sprintf("第 %d 天 | 内容: %s | 指纹: %s\n", 
        day, m.Context, hashString)
	    aa.WriteString(line)
		}

		fmt.Printf("奇迹发生了！")
}