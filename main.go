package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			termbox.SetCell(x, y, '@', termbox.ColorGreen, termbox.ColorGreen)
			termbox.Flush()
			time.Sleep(time.Millisecond * 50)
		}
	}
}
