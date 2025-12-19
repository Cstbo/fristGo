package main

import (
	"fmt"
	"os"
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
		fmt.Printf("今天是第%d天|日期%s|我今天%s\n ",day,m.date,m.Context)
	}
	 
	file,err := os.Create("love.txt")
	if err != nil{
		fmt.Println("Error..",err)
		return
	}
	defer file.Close()

	for day,m := range diary {
		line := fmt.Sprintf("今天是第%d天|日期%s|我今天%s\n ",day,m.date,m.Context)
	    file.WriteString(line)
		}

		fmt.Printf("奇迹发生了！")
}