package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/odsod/gb/cartridge"
)

func main() {
	romFile := flag.String("rom", "roms/tetris.gb", "The ROM to load")
	flag.Parse()

	rom, err := ioutil.ReadFile(*romFile)
	if err != nil {
		panic(err)
	}

	cart := cartridge.FromROM(rom)

	fmt.Printf("%#v\n", cart)
}
