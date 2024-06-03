package main

import (
	"fmt"

	"github.com/jamesyang124/celeritas"
)

func main() {
	res := celeritas.TestFunc1(1, 2)
	fmt.Println(res)
}
