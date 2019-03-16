package main

import (
	goled "github.com/talkkonnect/go-oled-i2c"
	"log"
)

var (
	OledText   string
	OledRow    int
	OledColumn int
)

func oledDisplay (OledClear bool,OledRow int, OledColumn int, OledText string){
	oled, err := goled.BeginOled()

	if err != nil {
		log.Fatal(err)
		return
	}
	defer oled.Close()

	if OledClear == true {
		oled.Clear()
		return
	}
	oled.SetCursor(OledRow, OledColumn)
	oled.Write(OledText)
}

func main() {

	oledDisplay(true,0,0,"") // clear the screen
	oledDisplay(false,0,1,"Hello 0 from go1111!\n")

}
