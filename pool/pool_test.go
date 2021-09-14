package pool

import (
	"fmt"
	"testing"
)

type Demo struct {
}

func (d *Demo) run() {
	fmt.Println("hello d")
}

func TestAddPool(t *testing.T) {
	p := &Pool{
		corePoolSize: 100,
	}
	p.Execute(&Demo{})
	t.Log("test finished")

}
