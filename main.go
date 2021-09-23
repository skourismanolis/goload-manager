package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
)

type Progress struct {
	textView *tview.TextView
	full     int
	limit    int
	progress chan int
}

// full is the maximum amount of value can be sent to channel
// limit is the progress bar size
func (p *Progress) Init(full int, limit int) chan int {
	p.progress = make(chan int)
	p.full = full
	p.limit = limit

	go func() { // Simple channel status gauge (progress bar)
		progress := 0
		for {
			progress += <-p.progress

			if progress > full {
				break
			}

			x := progress * limit / full
			p.textView.Clear()
			_, _ = fmt.Fprintf(p.textView, "channel status:  %s%s %d/%d",
				strings.Repeat("■", x),
				strings.Repeat("□", limit-x),
				progress, full)

		}
	}()

	return p.progress
}

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetChangedFunc(func() {
			app.Draw()
		})

	textView.SetBorder(true)
	// textView.SetBackgroundColor(tcell.ColorDefault)

	progress := Progress{textView: textView}
	progChan := progress.Init(360, 20)

	go func() { // update progress bar

		i := 0

		for {
			i++
			progChan <- 1

			if i > progress.full {
				close(progChan)
				break
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}
}
