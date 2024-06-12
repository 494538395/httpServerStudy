package demo

import (
	"fmt"
	"sync"
	"testing"
)

type person struct {
	name string
	age  int
}

var personPool = sync.Pool{New: func() interface{} { return new(person) }}

func TestSyncPool(t *testing.T) {

	for i := 0; i < 10; i++ {
		p1 := getPerson()
		fmt.Printf("p1：%p \n", p1)

		p2 := getPerson()
		fmt.Printf("p1：%p \n", p2)

		setPerson(p1)
		setPerson(p2)

	}

}

func getPerson() *person {
	return personPool.Get().(*person)
}

func setPerson(p *person) {
	// 重置对象状态
	p.name = ""
	p.age = 0
	personPool.Put(p)
}

func TestSlice(t *testing.T) {

}
