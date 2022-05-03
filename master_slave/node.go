/**
  @author: 黄睿楠
  @since: 2022/4/30
  @desc: 节点类
**/

package main

type Node struct {
	ID   string
	Port string
}

func NewNode(id, port string) *Node {
	node := new(Node)
	node.ID = id
	node.Port = port
	return node
}