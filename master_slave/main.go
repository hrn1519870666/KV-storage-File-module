/**
  @author: 黄睿楠
  @since: 2022/4/30
  @desc:
**/

package main

import (
	"log"
	"os"
	"time"
)

func main() {

	//定义三个节点  节点编号 - 端口号
	nodeTable = map[string]string{
		"A": ":9000",
		"B": ":9001",
		"C": ":9002",
	}

	//运行程序时，指定节点编号:node.exe A/B/C   os.Args[0]是node.exe
	if len(os.Args) < 1 {
		log.Fatal("程序参数不正确")
	}

	// A B C
	id := os.Args[1]
	// 创建raft节点实例
	node := NewNode(id, nodeTable[id])

	if id == "A" {
		go node.write()
	}

	if id == "B" {
		go node.read()
	}

	// 10分钟之后结束监听
	time.Sleep(600*time.Second)
}
