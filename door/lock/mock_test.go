// +build !rpi

package lock

func ExampleMock_Open() {
	m := Mock{}
	m.Open()
	// Output: Opened Lock
}
