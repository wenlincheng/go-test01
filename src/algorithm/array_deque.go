package main

import "errors"

/*
	ArrayDeque是一个双端队列实现，
	内部使用数组进行元素存储，不允许存储null值，
	可以高效的进行元素查找和尾部插入取出，
	是用作队列、双端队列、栈的绝佳选择，性能比LinkedList还要好。
*/

const (
	// 初始化最小容量
	MIN_INITIAL_CAPACITY int = 8
)

type ArrayDeque struct {
	elements []interface{}
	head     int
	tail     int
}

func (a *ArrayDeque) Add(e interface{}) bool {
	a.AddLast(e)
	return true
}

func (a *ArrayDeque) AddLast(e interface{}) {

}

// 将数组容量扩成之前的两倍
func (a *ArrayDeque) DoubleCapacity() error {
	p := a.head
	n := len(a.elements)
	r := n - p // number of elements to the right of p
	newCapacity := n << 1
	if newCapacity < 0 {
		return errors.New("Sorry, deque too big")
	}
	newArray := make([]interface{}, newCapacity)
	// 容量变为原来的两倍，然后把head之后的元素复制到新数组的开头，把剩余的元素复制到新数组之后
	a.arrayCopy(a.elements, p, newArray, 0, r)
	a.arrayCopy(a.elements, 0, newArray, r, p)
	a.elements = newArray
	a.head = 0
	a.tail = n

	return nil
}

// 复制数组
func (a *ArrayDeque) arrayCopy(src []interface{}, srcPos int, dest []interface{}, destPos int, length int) {
	for i := srcPos; i < srcPos+length; i++ {
		dest[destPos] = src[i]
		destPos++
	}
}
