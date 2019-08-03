// Generates sprite sheets & data used by AWO (outputs visuals.json & spritesheet.png)
package main

import "fmt"

func main() {
	test := unitFrameData{
		x: 4,
		y: 4,
		w: 4,
		h: 5,
	}

	fmt.Printf("%#v\n", test)
}
