package main

import "fmt"
import "sync"
import (
	"runtime"
)


var counter int = 0

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
func Count(lock *sync.Mutex){
	lock.Lock()
	counter ++
	Add(3,4)
	lock.Unlock()
}

func main(){
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10{
			break
		}

	}

}

