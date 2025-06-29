package bufferpool

import "sync"


type bpool struct{

	//16B buffer
	Pool16 *sync.Pool
	//64B buffer
	Pool64 *sync.Pool
	//1KB buffer
	Pool1k *sync.Pool
	//1MB buffer
	Pool1m *sync.Pool

	//page buffer
	PoolPage *sync.Pool

	len16 uint32
	len64 uint32
	len1k uint32
	len1m uint32

}

var (
	Buffer *bpool
)