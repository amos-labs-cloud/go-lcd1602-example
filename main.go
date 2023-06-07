package main

import (
	"github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/synchronized"
	"time"
)

func main() {
	lcdi := lcd1602.New(18, 23, []int{16, 12, 25, 24}, 16)
	syncedLCD := synchronized.NewSynchronizedLCD(lcdi)
	syncedLCD.Initialize()
	defer syncedLCD.Close()

	syncedLCD.WriteLines("Hello world!")
	time.Sleep(time.Second * 5)
	syncedLCD.WriteLines("")
	syncedLCD.WriteLines("from Amos Labs")
}
