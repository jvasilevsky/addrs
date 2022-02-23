package ipv6

import (
	"sync/atomic"
	"unsafe"
)

func swapTrieNodePtr(ptr **trieNode, old, new *trieNode) bool {
	return atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(
			unsafe.Pointer(ptr),
		),
		unsafe.Pointer(old),
		unsafe.Pointer(new),
	)
}
