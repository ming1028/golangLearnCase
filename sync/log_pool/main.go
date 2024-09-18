package main

import (
	"bytes"
	"os"
	"strconv"
	"sync"
)

const MaxBatchSize = 1000

var writePool = sync.Pool{
	New: func() any {
		return &BatchData{buffer: new(bytes.Buffer)}
	},
}

type BatchData struct {
	buffer *bytes.Buffer
	conter int
}

func getBatchData() *BatchData {
	v := writePool.Get()
	/*if v == nil {
		return &BatchData{
			buffer: new(bytes.Buffer),
		}
	}*/
	return v.(*BatchData)
}

func releaseBatchData(bd *BatchData) {
	bd.buffer.Reset()
	bd.conter = 0
	writePool.Put(bd)
}

func main() {
	file, _ := os.OpenFile("./test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	/*writePool.New = func() any {
		return &BatchData{buffer: new(bytes.Buffer)}
	}*/

	for i := 0; i < 1000; i++ {
		bd := getBatchData()
		for j := 0; j < MaxBatchSize; j++ {
			bd.buffer.WriteString("test data" + strconv.Itoa(i*MaxBatchSize) + "\n")
			bd.conter++
		}
		if bd.conter >= MaxBatchSize {
			file.Write(bd.buffer.Bytes())
			releaseBatchData(bd)
		}
	}
}
