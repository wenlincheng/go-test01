package employee

import (
	"fmt"
)

// 构造体
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

// 类似Java中的构造器
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}