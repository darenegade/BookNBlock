package main

import (
	"github.com/darenegade/BookNBlock/door/lock"
	"time"
)

func main() {
	lock.Lock.Open()
	time.Sleep(time.Second * 4)
}