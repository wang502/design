package main

import (
	"fmt"

	"github.com/wang502/design/transcoders"
)

func main() {
	cases := []string{
		"1.2.3.4",
		"2601:9:4f81:9700:803e:ca65:66e8:c21",
	}
	names := []string{
		"ip4",
		"ip6",
	}
	var p transcoders.Protocol

	for i, a := range cases {
		for _, pro := range transcoders.Protocols {
			if pro.Name == names[i] {
				p = pro
			}
		}
		fmt.Println(p.Transcoder.StringToByte(a))
	}
}
