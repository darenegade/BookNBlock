// +build !rpi

package driver

import (
	"fmt"
)

type (
	Mock struct{}
)

func (Mock) Open() {
	fmt.Println("Opened Lock")
}
