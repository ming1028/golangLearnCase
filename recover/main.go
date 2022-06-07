package main

import "log"

func main() {

}

func protect(g func()) {
	defer func() {
		log.Println("panic")
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
}
