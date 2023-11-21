package pool

import (
	"bytes"
	"os"
	"sync"
	"testing"
)

var pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

var fileName = "sync_pool.log"
var data = make([]byte, 10000)

func main() {
}

func BenchmarkWriteFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		buf.Reset()
		buf.Write(data)
		os.WriteFile(fileName, buf.Bytes(), 0644)
	}
}

func BenchmarkWriteFileWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := pool.Get().(*bytes.Buffer)
		buf.Reset()
		buf.Write(data)
		os.WriteFile(fileName, buf.Bytes(), 0644)
		pool.Put(buf)
	}
}
