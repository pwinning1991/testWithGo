package parallel

import (
	"fmt"
	"testing"
	"time"
)

//func TestSomething(t *testing.T) {
//	t.Parallel()
//
//}

//func TestA(t *testing.T) {
//	t.Parallel()
//	time.Sleep(time.Second)
//}

func TestB(t *testing.T) {
	fmt.Println("setup")
	defer fmt.Println("deferred teardown")
	t.Run("group", func(t *testing.T) {
		t.Run("sub1", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("sub1 done")
		})
		t.Run("sub2", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("sub2 done")
		})
	})
	fmt.Println("teardown")
}

func TestGotcha(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("i=%d", i), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with i=%d", i)
		})
	}
}
