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

	// Start creating events.
	e := make(chan termbox.Event)
	go func() {
		for {
			e <- termbox.PollEvent()
		}
	}()

	// Step through the x axis (left to right).
	for x := 0; x < w; x++ {
		// Step through the y axis (top to bottom).
		for y := 0; y < h; y++ {
			select {
			// If we've pressed escape.
			case ev := <-e:
				if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
					// quit!
					return
				}
			default:
				// Draw to the terminal.
				termbox.SetCell(x, y, '@', termbox.ColorGreen, termbox.ColorGreen)
				termbox.Flush()
				time.Sleep(time.Millisecond * 5)
			}
		}
	}
}
