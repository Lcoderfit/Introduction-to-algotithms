package main

import (
	"container/list"
	"fmt"
)

func main() {
	p := list.New()
	for i := 0; i < 5; i++ {
		p.PushBack(i)
	}

	for e := p.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(p.Front().Value, p.Back().Value)
	p.InsertAfter(10, p.Front())
	for e := p.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	p.MoveToFront(p.Back())
	for e := p.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l2 := list.New()
	l2.PushBackList(p)
	fmt.Println(l2.Len())

	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// func (e *Element) Next() *Element  //返回该元素的下一个元素，如果没有下一个元素则返回nil
// func (e *Element) Prev() *Element//返回该元素的前一个元素，如果没有前一个元素则返回nil。
// type List
// func New() *List //返回一个初始化的list
// func (l *List) Back() *Element //获取list l的最后一个元素
// func (l *List) Front() *Element //获取list l的第一个元素
// func (l *List) Init() *List  //list l初始化或者清除list l
// func (l *List) InsertAfter(v interface{}, mark *Element) *Element  //在list l中元素mark之后插入一个值为v的元素，并返回该元素，如果mark不是list中元素，则list不改变。
// func (l *List) InsertBefore(v interface{}, mark *Element) *Element//在list l中元素mark之前插入一个值为v的元素，并返回该元素，如果mark不是list中元素，则list不改变。
// func (l *List) Len() int //获取list l的长度
// func (l *List) MoveAfter(e, mark *Element)  //将元素e移动到元素mark之后，如果元素e或者mark不属于list l，或者e==mark，则list l不改变。
// func (l *List) MoveBefore(e, mark *Element)//将元素e移动到元素mark之前，如果元素e或者mark不属于list l，或者e==mark，则list l不改变。
// func (l *List) MoveToBack(e *Element)//将元素e移动到list l的末尾，如果e不属于list l，则list不改变。
// func (l *List) MoveToFront(e *Element)//将元素e移动到list l的首部，如果e不属于list l，则list不改变。
// func (l *List) PushBack(v interface{}) *Element//在list l的末尾插入值为v的元素，并返回该元素。
// func (l *List) PushBackList(other *List)//在list l的尾部插入另外一个list，其中l和other可以相等。
// func (l *List) PushFront(v interface{}) *Element//在list l的首部插入值为v的元素，并返回该元素。
// func (l *List) PushFrontList(other *List)//在list l的首部插入另外一个list，其中l和other可以相等。
// func (l *List) Remove(e *Element) interface{}//如果元素e属于list l，将其从list中删除，并返回元素e的值。
