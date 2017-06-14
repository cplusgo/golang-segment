package main

import (
	"github.com/cplusgo/golang-segment"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

//加载词库和分词的时间消耗均已打印出来,时间单位是微妙
func main() {
	tireTree := golang_segment.NewTireTree()
	start := time.Now().Nanosecond()
	tireTree.LoadDict("data/words.dict")
	dut := (time.Now().Nanosecond() - start) / 1000
	log.Println("加载字库时间消耗:", dut)

	var sentence string = "作为迄今为止世界出版史上最高发行量的字典，《新华字典》自从上世纪50年代开始出版以来，伴随了几代中国人的启蒙教育，至今仍广为社会大众所接受。"
	start = time.Now().Nanosecond()
	words := tireTree.Search(sentence)
	dut = (time.Now().Nanosecond() - start) / 1000
	log.Println("分词时间消耗:", dut)
	log.Println(words)
}
