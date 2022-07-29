package main
//单链接表
// node：包含一个数据域，一个指针域(指向下一个节点)
// LList：包含头指针(指向第一个节点)，链表长度

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LList struct {
	Header *Node
	Length int64
}

func CreateNode(v interface{}) *Node {
	return &Node{v, nil}
}
func CreateLList() *LList {
	header := CreateNode(nil)
	return &LList{header, 0}
}
func (l *LList) Add(data interface{}){
	node := CreateNode(data)
	defer func ()  {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = node
	} else {
		node.Next = l.Header
		l.Header = node// 头插法
	}
}

// 往链表的表尾添加一个节点
func (l *LList) Append(data interface{}) {
	node := CreateNode(data)
	defer func ()  {
		l.Length++
	}()
	if l.Length == 0 {
		l.Header = node
	} else {
		nextNode := l.Header
		for  nextNode.Next != nil { //循环找到最后一个节点
			nextNode = nextNode.Next
		} 
		nextNode.Next = node //把新节点地址给最后一个节点的Next， 引用操作，所以直接赋值就好
	}
}

// 往i后面插入一个元素
func (l *LList) Insert(i int64, data interface{})  {

	defer func ()  {
		l.Length++
	}()

	if i >=l.Length {
		l.Append(data)
		return
	}
	node := CreateNode(data)
	nextNode := l.Header
	var j int64 = 0
	for ; j < i; j++ { //找到第i个节点
		nextNode = nextNode.Next  
	}
	node.Next = nextNode.Next
	nextNode.Next = node
}

func (l *LList) Delete(i int64)  {
	if i > l.Length {
		return
	}
	nextNode := l.Header
	var j int64 = 0
	for ; j < i-1; j++ {
		nextNode = nextNode.Next 
	}
	nextNode.Next = nextNode.Next.Next
}


func (l *LList) Scan()  {
	node := l.Header
	if node == nil {
		return
	}
	for node != nil {
		fmt.Println("value is ： %v",node.Data)
		node = node.Next
	}
	return
}
func main()  {
	Header := CreateLList()
	Header.Add(1)
	Header.Add(1)
	Header.Delete(1)
	Header.Insert(0,9)
	Header.Add(6)
	Header.Add(4)
	Header.Add(3)
	Header.Append(10)
	Header.Scan()
}
