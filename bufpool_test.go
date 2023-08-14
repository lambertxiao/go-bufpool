package mempool

import (
	"testing"
	"time"
)

func TestGoBufPoolGet(t *testing.T) {
	type Foo struct{}
	bufPool := NewBufPool(5, 1024)

	capacity := bufPool.Cap()
	if capacity != 5 {
		t.Errorf("The capacity of the memory pool is incorrect")
	}

	for i := 0; i < 5; i++ {
		tmpBuf := bufPool.Get()
		if tmpBuf == nil {
			t.Fatal("Failed to retrieve a node from the pool")
		}
	}

	bufPool.Destory()

	if !bufPool.stack.IsEmpty() {
		t.Errorf("The memory pool has not been emptied")
	}
}

func BenchmarkGet(b *testing.B) {
	bufPool := NewBufPool(10000, 1024*1024)
	for i := 0; i < 10000; i++ {
		bufPool.Get()
	}

	b.ReportAllocs()
}

func BenchmarkGetByTime(b *testing.B) {
	bufPool := NewBufPool(10000, 1024)
	for i := 0; i < 10001; i++ {
		bufPool.GetByTime(2 * time.Second)
	}

	b.ReportAllocs()
}
