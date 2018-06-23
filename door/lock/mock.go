// +build !rpi

package lock

import (
	"fmt"
)

type (
	Mock struct{}
)

var Lock = Mock{}

func (Mock) Open() {
	fmt.Println("Opened Lock")
}
