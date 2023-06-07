package main

import (
	"github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/synchronized"
	"log"
	"time"
)

func main() {
	lcdi := lcd1602.New(18, 23, []int{16, 12, 25, 24}, 16)
	syncedLCD := synchronized.NewSynchronizedLCD(lcdi)
	syncedLCD.Initialize()
	defer syncedLCD.Close()

	log.Println("Resetting screen")
	syncedLCD.WriteLines()

	log.Println("Printing a message")
	syncedLCD.WriteLines("Hello world!")
	log.Println("Sleeping for 5 seconds")
	time.Sleep(time.Second * 5)
	log.Println("Resetting the screen")
	syncedLCD.WriteLines()
	log.Println("Writing another message")
	syncedLCD.WriteLines("from Amos Labs")
	time.Sleep(time.Second * 2)
	log.Println("Resetting the screen")
	syncedLCD.WriteLines()
	log.Println("Writing multiple lines")
	syncedLCD.WriteLines("This is a ", "test of the lcd")
}
