package demo

/*
	测试框架 testing
	1、测试文件 demo_test.go 的文件名后缀必须为 _test，文件名前半部分可随意取。
	2、测试函数 Test0001,Test0002等，函数名必须以 Test 前缀，函数名后半部分可随意取，参数必须为 (test *testing.T)
*/
import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	fmt.Println("Demo test Run")

	fmt.Println("Hello World")
}

func Test02(t *testing.T) {
	fmt.Println("testing run")
}

func TestAbs(t *testing.T) {
	result := math.Abs(-1)
	if result != 1 {
		t.Errorf("Abs(-1) = %f; want 1", result)
	}

}

/*
	基准测试
*/
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello.txt")
	}
}

func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) { fmt.Println("Hello") })
	t.Run("A=2", func(t *testing.T) { fmt.Println("Hello") })
	t.Run("B=1", func(t *testing.T) { fmt.Println("Hello") })
	// <tear-down code>
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	// setup
	code := m.Run()
	// teardown
	os.Exit(code)
}
