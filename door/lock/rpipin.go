// +build rpi

package lock

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

type (
	Pin struct {
		pin gpio.PinOut
	}
)

func init() {
	_, err := host.Init()
	if err != nil {
		panic(err)
	}

	PinLock.pin.Out(gpio.Low)
}

var Lock = Pin{
	pin: rpi.P1_38,
}

func (p Pin) Open() {
	p.pin.Out(gpio.High)
	time.AfterFunc(time.Second*2, func() {
		p.pin.Out(gpio.Low)
	})
}
