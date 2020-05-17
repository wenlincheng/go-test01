package employee

import (
	"fmt"
	"github.com/wenlincheng/go-common/collection"
	"testing"
)

func TestEmployee(t *testing.T) {
	e := New("Wen", "Lincheng", 3, 24)
	e.LeavesRemaining()
	deque := collection.ArrayDeque{}

	deque.Add(1)

	_, i := deque.GetLast()

	fmt.Println(i)

}
