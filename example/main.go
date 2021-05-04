package main

import (
	"log"

	"github.com/c93614/opencc-go"
)

func main() {
	converter, err := opencc.NewConverter("s2twp.json")
	log.Print(converter, err)
	result := converter.Convert([]byte("制定接口规范，编写核心代码，推动自动化测试和部署，重构和优化现有技术架构；"))
	converter.Close()
	log.Println(string(result))
}
