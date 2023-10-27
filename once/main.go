package main

import (
	"fmt"
	"sync"
)

func main() {
	tOnce := sync.Once{}
	fmt.Println(tOnce)
}

func getToken() {

}
