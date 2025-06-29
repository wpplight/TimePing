package bufferpool

import (
	"sync"
	"sync/atomic"
	"timeping/pkg/upool"
)


func (b *bpool)Init() {
	b.Pool16=&sync.Pool{
		New: func() any {
			return b.make16B()
		},
	}
	b.Pool1k=&sync.Pool{
		New: func() any {
			return b.make1K()
		},
	}
	b.Pool1m=&sync.Pool{
		New: func() any {
			return b.make1m()
		},
	}
	b.Pool64=&sync.Pool{
		New: func() any {
			return b.make64B()
		},
	}
	upool.Init_upool()
}

func (b *bpool)make64B() []byte {
	atomic.AddUint32(&b.len64,1)
	return make([]byte, 64)
}

func (b *bpool)make16B() []byte {
	atomic.AddUint32(&b.len16,1)
	return make([]byte ,16)
}

func (b *bpool)make1K() []byte {
	atomic.AddUint32(&b.len1k,1)
	return make([]byte,1<<10)
}

func (b *bpool)make1m() []byte {
	atomic.AddUint32(&b.len1m,1)
	return make([]byte,1<<20)
}