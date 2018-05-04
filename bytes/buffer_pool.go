package bytes

import (
	"bytes"
	"sync"
)

// TextBufferPool text buffer pool
var TextBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 16<<10)) // 16KB
	},
}

// MediaBufferPool media buffer pool
var MediaBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 10<<20)) // 10MB
	},
}
