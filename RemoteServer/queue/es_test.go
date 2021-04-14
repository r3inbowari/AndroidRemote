package queue

import "testing"

func TestESQ(t *testing.T) {
	es := NewQueue(100)
	es.Put(123)
	es.Put(1232)
	for i := 0; i < 100; i++ {
		es.Put(12)
	}
	val, ok, q := es.Get()
	println(q)
	println(ok)
	println(val.(int))
}
