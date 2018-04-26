// +build !rpi

package driver

func ExampleMock_Open() {
	m := Mock{}
	m.Open()
	// Output: Opened Lock
}
