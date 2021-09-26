package main

import (
	// "fmt"
	// "strings"

	"fmt"
	"sync"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/gosuri/uiprogress"
	// "github.com/skourismanolis/goload-manager/test"
)

func initBar(resp *grab.Response) *uiprogress.Bar {
	bar := uiprogress.AddBar(100) // Add a new bar
	// optionally, append and prepend completion and elapsed time
	bar.Empty = '_'
	bar.Fill = '#'
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return resp.Filename
	})
	bar.AppendCompleted()
	// bar.PrependElapsed()
	return bar
}

func init() {
	uiprogress.Start() // start rendering
}

func main() {
	var wg sync.WaitGroup

	client := grab.NewClient()

	// downloads := make([]int, 0)

	for {
		var line string
		fmt.Printf("URL> ")
		fmt.Scanln(&line)

		req, _ := grab.NewRequest("", line)
		download := client.Do(req)

		wg.Add(1)
		go func() {
			defer wg.Done()
			bar := initBar(download)
			t := time.NewTicker(time.Millisecond * 500)
			defer t.Stop()

		Loop:
			for {
				select {
				case <-t.C:
					bar.Set(int(download.Progress() * 100))
				case <-download.Done:
					// download is complete
					break Loop
				}
			}
		}()
	}

	wg.Wait()
}
