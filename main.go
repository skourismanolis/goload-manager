package main

import (
	// "fmt"
	// "strings"

	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/gosuri/uiprogress"
	// "github.com/skourismanolis/goload-manager/test"
)

func initBar(download *grab.Response) *uiprogress.Bar {
	bar := uiprogress.AddBar(100) // Add a new bar
	// optionally, append and prepend completion and elapsed time
	bar.Empty = '_'
	bar.Fill = '#'
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return download.Filename
	})
	bar.AppendCompleted()
	// bar.PrependElapsed()
	return bar
}

func monitorDownload(download *grab.Response) {
	bar := initBar(download)

	fmt.Println("asd")
	t := time.NewTicker(time.Millisecond * 50)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Println(download.Progress())
			bar.Set(int(download.Progress() * 100))
		case <-download.Done:
			// download is complete
			break Loop
		}
	}

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
		if strings.ToLower(line) == "exit" {
			break
		}

		req, _ := grab.NewRequest("./downloads", line)

		wg.Add(1)
		go func() {
			download := client.Do(req)
			defer wg.Done()
			monitorDownload(download)
		}()

	}

	wg.Wait()
}
