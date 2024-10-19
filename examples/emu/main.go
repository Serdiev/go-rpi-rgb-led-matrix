package main

import (
	"fmt"
	"image/color"

	"github.com/Serdiev/go-rpi-rgb-led-matrix/emulator"
)

func main() {
	h := 20
	w := 20
	e := emulator.NewEmulator(w, h, 0, false)

	for i := 0; i <= (h*w)-1; i++ {
		e.Set(i, color.RGBA{R: 0, G: 255, B: 0, A: 255})
		fmt.Println(i)
	}

	e.Init()

	// leds := []color.Color{
	// 	color.RGBA{R: 0, G: 255, B: 0, A: 255}, // Green
	// 	color.RGBA{R: 255, G: 0, B: 0, A: 255}, // Red
	// 	color.RGBA{R: 0, G: 0, B: 255, A: 255}, // Blue
	// }
	err := e.Render()
	fmt.Println(err)

}
