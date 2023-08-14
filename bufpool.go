package mempool

import (
	"arena"
	"time"
)

type GoBufPool struct {
	fn    func() interface{}
	stack *Stack
	mem   *arena.Arena
}

func NewBufPool(capacity int, itemSize int) *GoBufPool {
	mem := arena.NewArena()
	stack := newStack()

	for i := 0; i < capacity; i++ {
		paritalBuf := arena.MakeSlice[byte](mem, itemSize, itemSize)
		n := newNode(paritalBuf)
		stack.Push(n)
	}

	return &GoBufPool{
		stack: stack,
		mem:   mem,
	}
}

func (g *GoBufPool) Get() []byte {
	for {
		node, err := g.stack.Pop()
		if err != nil {
			continue
		}
		return node.Value.([]byte)
	}
}

func (g *GoBufPool) GetByTime(timeout time.Duration) []byte {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return nil // 超时返回空值
		default:
			node, err := g.stack.Pop()
			if err != nil {
				continue
			}
			return node.Value.([]byte)
		}
	}
}

func (g *GoBufPool) Put(data []byte) {
	n := newNode(data)
	g.stack.Push(n)
}

func (g *GoBufPool) Cap() int {
	return g.stack.Cap()
}

func (g *GoBufPool) Destory() {
	for !g.stack.IsEmpty() {
		g.stack.Pop()
	}
	g.mem.Free()
}
