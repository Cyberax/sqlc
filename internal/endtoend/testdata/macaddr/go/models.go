// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"net"
)

type Foo struct {
	Bar  bool
	Addr net.HardwareAddr
}

func (t *Foo) GetBar() bool {
	return t.Bar
}
func (t *Foo) GetAddr() net.HardwareAddr {
	return t.Addr
}
