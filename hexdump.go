// dumps hex strings from stdout
// $ hexdump 202122232425262728292A2B2C2D2E2F303132333435363738393A3B3C3D3E3F
// 00000000  20 21 22 23 24 25 26 27  28 29 2a 2b 2c 2d 2e 2f  | !"#$%&'()*+,-./|
// 00000010  30 31 32 33 34 35 36 37  38 39 3a 3b 3c 3d 3e 3f  |0123456789:;<=>?|
//
package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	for _, a := range os.Args[1:] {
		bytes, err := hex.DecodeString(a)
		if err != nil {
			continue
		}
		fmt.Printf("%s", hex.Dump(bytes))
	}
}
