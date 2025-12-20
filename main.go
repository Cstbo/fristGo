package main //声明主程序

import (
	"crypto/sha256" //碎纸机
	"encoding/hex"  //16进制转换
	"fmt"
	"os" //引入操作系统工具包，用来创建和保存文件
)

type Memory struct { //定义了一个Memory结构体
	date    string
	Context string
}

func main() {
	diary := make(map[int]Memory) //创建了一个日记本

	diary[1] = Memory{date: "12/19", Context: "喜欢你"}
	diary[2] = Memory{date: "12/20", Context: "依旧喜欢你"}
	diary[3] = Memory{date: "12/21", Context: "还是很喜欢你"}

	fmt.Println("今天的状态是:")
	for day, m := range diary {

		record := fmt.Sprintf("%d%s%s", day, m.date, m.Context) //把所有内容糅合到一起，定义为record
		h := sha256.New()                                       //拿出碎纸机
		h.Write([]byte(record))                                 //把糅合的内容切成字节

		hash := h.Sum(nil)                     //拿到二进制哈希
		hashString := hex.EncodeToString(hash) //翻译成人类能看懂的

		fmt.Printf("\033[35m第 %d 天\033[0m | 内容: %s | \033[32m指纹: %s\033[0m\n",
			day, m.Context, hashString) //炫彩打印
	}
	//写入文件
	aa, err := os.Create("love.txt") //创建一个love文本文件，err检查是否出错
	if err != nil {                  //如果创建失败
		fmt.Println("Error..", err) //报错
		return                      //退出程序
	}
	defer aa.Close() //在最后关闭文件

	for day, m := range diary {

		record := fmt.Sprintf("%d%s%s", day, m.date, m.Context)
		g := sha256.New()
		g.Write([]byte(record))
		hash := g.Sum(nil)
		hashString := hex.EncodeToString(hash)

		line := fmt.Sprintf("第 %d 天 | 内容: %s | 指纹: %s\n", //排版
			day, m.Context, hashString)
		aa.WriteString(line) //通过aa这个遥控器，将line写入love文本文件
	}

	fmt.Printf("奇迹发生了！") //结束
}
