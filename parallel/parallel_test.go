package parallel

import (
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	t.Parallel()

}

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestB(t *testing.T) {}
