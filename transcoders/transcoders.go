package transcoders

import (
	"fmt"
	"net"
)

// Transcoder ...
type Transcoder interface {
	StringToByte(string) ([]byte, error)
	BytesToString([]byte) (string, error)
}

// NewTranscoderFromFunctions ...
func NewTranscoderFromFunctions(s2b func(string) ([]byte, error), b2s func([]byte) (string, error)) Transcoder {
	return twrp{s2b, b2s}
}

type twrp struct {
	strtobyte func(string) ([]byte, error)
	bytetostr func([]byte) (string, error)
}

func (t twrp) StringToByte(s string) ([]byte, error) {
	return t.strtobyte(s)
}

func (t twrp) BytesToString(b []byte) (string, error) {
	return t.bytetostr(b)
}

var TranscoderIP4 = NewTranscoderFromFunctions(ip4StB, ipBtS)
var TranscoderIP6 = NewTranscoderFromFunctions(ip6StB, ipBtS)

func ip4StB(s string) ([]byte, error) {
	i := net.ParseIP(s).To4()
	if i == nil {
		return nil, fmt.Errorf("failed to parse ip4 addr: %s", s)
	}
	return i, nil
}

func ip6StB(s string) ([]byte, error) {
	i := net.ParseIP(s).To16()
	if i == nil {
		return nil, fmt.Errorf("failed to parse ip6 addr: %s", s)
	}
	return i, nil
}

func ipBtS(b []byte) (string, error) {
	return net.IP(b).String(), nil
}
