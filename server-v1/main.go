package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入: ")
	var text, _ = reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	fmt.Println("打印去除换行的文本：" + text)
	fmt.Println("打印split后的切片：", strings.Split(text, " "))
}
