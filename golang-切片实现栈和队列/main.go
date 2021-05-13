package main
import (
	"fmt"
)
// goalng实现栈

type Stack struct {
	List []interface{}
}

func NewStack() *Stack {
	return new(Stack)
} 

func (s *Stack) Push(data interface{})  {
	s.List = append(s.List, data)
	return
}
func (s *Stack) Pop() interface{} {
	len := len(s.List)
	var value interface{}
	if len == 0 {
		return nil
	} else if len == 1 {
		value = s.List[0]
		s.List = []interface{}{}
	} else {
		value = s.List[len-1]
		s.List = s.List[0:len-1]
	}
	return value
}

func (s *Stack) Print() {
    for _,v := range s.List {
        fmt.Printf("%v  ",v)
    }
	fmt.Printf("\n")
}
func main(){
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
	stack.Print()
    res1:= stack.Pop()
    res2:= stack.Pop()
    res3:= stack.Pop()
	stack.Print()
    fmt.Println(res1,res2,res3)



	que := NewQueue()
    que.Push(1)
    que.Push(2)
	que.Print()
    que1:= que.Pop()
    que2:= que.Pop()
    que3:= que.Pop()
	stack.Print()
    fmt.Println(que1,que2,que3)
}
// golang实现队列


type Queue struct {
	List []interface{}
}
func NewQueue() *Queue  {
	list := make([]interface{},0)
	return &Queue{List:list}
}

func (q *Queue)Push(data interface{}) {
	q.List = append(q.List,data)
	return
}
func (q *Queue)Pop() interface{} {
	len := len(q.List)

	if len == 0 {
		return nil
	}
	res := q.List[0]
	q.List = q.List[1:] 
	return res
}

func (q *Queue)Print(){
    for _,val := range q.List {
        fmt.Println(val)
    }
}