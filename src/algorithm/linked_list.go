package main

import (
	"errors"
	"fmt"
)

var (
	nodeNotExit = errors.New("元素不存在")
)

// 双向链表
type LinkedList struct {
	size     int   // 链表长度
	modCount int   // 修改次数
	first    *Node // 头节点
	last     *Node // 尾节点
}

// 节点
type Node struct {
	item interface{} // 值
	prev *Node       // 前一个结点指针
	next *Node       // 下一个结点指针
}

// 链表的大小
func (list *LinkedList) Size() int {
	return list.size
}

// 是否为空
func (list *LinkedList) IsEmpty() bool {
	if list.size > 0 {
		return false
	}
	return true
}

// 判断链表是否包含某值
func (list *LinkedList) Contains(item interface{}) bool {
	return list.indexOf(item) != -1
}

// 该列表中第一次出现的指定元素的索引 链表不包含该元素返回-1
func (list *LinkedList) indexOf(item interface{}) int {
	index := 0
	if item == nil {
		for x := list.first; x != nil; x = x.next {
			if x.item == nil {
				return index
			}
			index++
		}
	} else {
		for x := list.first; x != nil; x = x.next {
			if item == x.item {
				return index
			}
			index++
		}
	}
	return -1
}

// 该列表中最后一次出现的指定元素的索引 链表不包含该元素返回-1
func (list *LinkedList) lastIndexOf(item interface{}) int {
	index := 0
	if item == nil {
		for x := list.last; x != nil; x = x.prev {
			index--
			if x.item == nil {
				return index
			}
		}
	} else {
		for x := list.last; x != nil; x = x.prev {
			index--
			if item == x.item {
				return index
			}
			index++
		}
	}
	return -1
}

// 在链表添加元素 默认在尾部添加
func (list *LinkedList) Add(item interface{}) bool {
	list.linkLast(item)
	return true
}

// 在链表尾部添加元素
func (list *LinkedList) AddLast(item interface{}) {
	list.linkLast(item)
}

// 在链表头部添加元素
func (list *LinkedList) AddFirst(item interface{}) {
	list.linkFirst(item)
}

// 创建一个新的节点
func (list *LinkedList) newNode(prev *Node, item interface{}, next *Node) *Node {
	n := Node{
		item: item,
		prev: prev,
		next: next,
	}
	return &n
}

// 在链表最前添加节点
func (list *LinkedList) linkFirst(e interface{}) {
	f := list.first
	newNode := list.newNode(nil, e, f)
	list.first = newNode
	if f == nil {
		list.last = newNode
	} else {
		f.prev = newNode
	}
	list.size++
	list.modCount++
}

// 在链表最后添加节点
func (list *LinkedList) linkLast(e interface{}) {
	l := list.last
	newNode := list.newNode(l, e, nil)
	list.last = newNode
	if l == nil {
		list.first = newNode
	} else {
		l.next = newNode
	}
	list.size++
	list.modCount++
}

// 将值添加到指定节点前
func (list *LinkedList) linkBefore(e interface{}, succ *Node) {
	pred := succ.prev
	newNode := list.newNode(pred, e, succ)
	succ.prev = newNode
	if pred == nil {
		list.first = newNode
	} else {
		pred.next = newNode
	}
	list.size++
	list.modCount++
}

func (list *LinkedList) isElementIndex(index int) bool {
	return index >= 0 && index < list.size
}

// 检查索引是否存在
func (list *LinkedList) checkElementIndex(index int) error {
	if !list.isElementIndex(index) {
		return list.outOfBoundsMsg(index)
	}
	return nil
}

func (list *LinkedList) checkPositionIndex(index int) error {
	if !list.isPositionIndex(index) {
		return list.outOfBoundsMsg(index)
	}
	return nil
}

func (list *LinkedList) outOfBoundsMsg(index int) error {
	s := fmt.Sprintf("IndexOutOfBounds Index: %d, Size: %d", index, list.size)
	return errors.New(s)
}

func (list *LinkedList) isPositionIndex(index int) bool {
	return index >= 0 && index <= list.size
}

// 删除指定索引的节点
func (list *LinkedList) RemoveIndex(index int) (error, interface{}) {
	err := list.checkElementIndex(index)
	if err != nil {
		return err, nil
	}
	return nil, list.unlink(list.node(index))

}

// 删除某个值
func (list *LinkedList) Remove(item interface{}) bool {
	if item == nil {
		for x := list.first; x != nil; x = x.next {
			if x.item == nil {
				list.unlink(x)
				return true
			}
		}
	} else {
		for x := list.first; x != nil; x = x.next {
			if item == x.item {
				list.unlink(x)
				return true
			}
		}
	}
	return false
}

// 删除第一个节点
func (list *LinkedList) RemoveFirst() (error, interface{}) {
	f := list.first
	if f == nil {
		return nodeNotExit, nil
	}
	return nil, list.unlinkFirst(f)
}

// 删除最后一个节点
func (list *LinkedList) RemoveLast() (error, interface{}) {
	l := list.last
	if l == nil {
		return nodeNotExit, nil
	}
	return nil, list.unlinkLast(l)
}

// 取消链接指定节点
func (list *LinkedList) unlink(x *Node) interface{} {
	element := x.item
	next := x.next
	prev := x.prev

	if prev == nil {
		list.first = next
	} else {
		prev.next = next
		x.prev = nil
	}

	if next == nil {
		list.last = prev
	} else {
		next.prev = prev
		x.next = nil
	}

	x.item = nil
	list.size--
	list.modCount++
	return element
}

// 取消链接最后一个节点
func (list *LinkedList) unlinkLast(l *Node) interface{} {
	element := l.item
	prev := l.prev
	l.item = nil
	l.prev = nil
	list.last = prev
	if prev == nil {
		list.first = nil
	} else {
		prev.next = nil
	}
	list.size--
	list.modCount++
	return element
}

// 取消链接第一个节点
func (list *LinkedList) unlinkFirst(f *Node) interface{} {
	element := f.item
	next := f.next
	f.item = nil
	f.prev = nil
	list.first = next
	if next == nil {
		list.last = nil
	} else {
		next.prev = nil
	}
	list.size--
	list.modCount++
	return element
}

// 返回指定索引的节点
func (list *LinkedList) node(index int) *Node {
	if index < (list.size >> 1) {
		x := list.first
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x
	} else {
		x := list.last
		for i := list.size - 1; i > index; i-- {
			x = x.prev
		}
		return x
	}
}

// 获取第一个节点
func (list *LinkedList) GetFirst() (error, interface{}) {
	f := list.first
	if f == nil {
		return errors.New("元素不存在"), nil
	}
	return nil, f.item
}

// 获取最后一个节点
func (list *LinkedList) GetLast() (error, interface{}) {
	l := list.last
	if l == nil {
		return errors.New("元素不存在"), nil
	}
	return nil, l.item
}

// 清空链表
func (list *LinkedList) Clear() {
	for x := list.first; x != nil; {
		next := x.next
		x.item = nil
		x.next = nil
		x.prev = nil
		x = next
	}
	list.first = nil
	list.last = nil
	list.size = 0
	list.modCount++
}

// 根据索引获取值
func (list *LinkedList) Get(index int) (error, interface{}) {
	err := list.checkElementIndex(index)
	if err != nil {
		return err, nil
	}
	return nil, list.node(index).item
}

// 替换指定索引的值
func (list *LinkedList) Set(index int, item interface{}) (error, interface{}) {
	err := list.checkElementIndex(index)
	if err != nil {
		return err, nil
	}
	x := list.node(index)
	oldVal := x.item
	x.item = item
	return nil, oldVal
}

// 将值添加在指定索引位置
func (list *LinkedList) addIndex(index int, item interface{}) error {
	err := list.checkPositionIndex(index)
	if err != nil {
		return err
	}
	if index == list.size {
		list.linkLast(item)
	} else {
		list.linkBefore(item, list.node(index))
	}
	return nil
}

// 检索但不删除此列表的头（第一个元素）
func (list *LinkedList) Element() (error, interface{}) {
	return list.GetFirst()
}

// 检索并删除此列表的头（第一个元素）
func (list *LinkedList) Poll() interface{} {
	f := list.first
	if f == nil {
		return nil
	}
	return list.unlinkFirst(f)
}

// 检索并删除此列表的第一个元素，如果此列表为空，则返回nil
func (list *LinkedList) PollFirst() interface{} {
	f := list.first
	if f == nil {
		return nil
	}
	return list.unlinkFirst(f)
}

// 检索并删除此列表的第一个元素，如果此列表为空，则返回nil
func (list *LinkedList) PollLast() interface{} {
	l := list.last
	if l == nil {
		return nil
	}
	return list.unlinkLast(l)
}

// 将元素插入此列表的开头
func (list *LinkedList) Push(item interface{}) {
	list.AddFirst(item)
}

// 将元素插入此列表的开头
func (list *LinkedList) Pop() (error, interface{}) {
	err, item := list.RemoveFirst()
	if err != nil {
		return err, nil
	}
	return nil, item
}