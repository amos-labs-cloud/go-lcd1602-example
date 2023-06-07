package main

import (
	"github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/synchronized"
)

func main() {
	lcdi := lcd1602.New(18, 23, []int{24, 25, 12, 16}, 16)
	syncedLCD := synchronized.NewSynchronizedLCD(lcdi)
	syncedLCD.Initialize()
	defer syncedLCD.Close()

	syncedLCD.WriteLines("Hello world!")
}
