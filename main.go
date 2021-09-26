package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/gosuri/uiprogress"
	"github.com/inhies/go-bytesize"
	// "github.com/skourismanolis/goload-manager/test"
)

func initBar(download *grab.Response) *uiprogress.Bar {
	bar := uiprogress.AddBar(100) // Add a new bar
	// optionally, append and prepend completion and elapsed time
	bar.Empty = '_'
	bar.Fill = '#'
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		tmp := strings.Split(download.Filename, "/")
		return tmp[len(tmp)-1]
	})
	bar.AppendFunc(func(b *uiprogress.Bar) string {
		eta := ""
		if time.Until(download.ETA()) >= 0 {
			eta = time.Until(download.ETA()).Round(time.Second).String()
		}

		progress := fmt.Sprintf("%.1f", download.Progress()*100)
		rate := bytesize.New(download.BytesPerSecond()).String() + "/s"
		return progress + "% " + rate + " " + eta
	})
	// bar.AppendCompleted()
	// bar.PrependElapsed()
	return bar
}

func monitorDownload(download *grab.Response, bar *uiprogress.Bar) {
	t := time.NewTicker(time.Millisecond * 50)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			bar.Set(int(download.Progress() * 100))
		case <-download.Done:
			bar.Set(100)
			// download is complete
			break Loop
		}
	}

}

// func init() {
// }

func main() {
	uiprogress.Start() // start rendering
	var wg sync.WaitGroup

	client := grab.NewClient()

	// downloads := make([]int, 0)

	// for {
	var line string
	line = "https://sabnzbd.org/tests/internetspeed/50MB.bin"
	// line = "https://golang.org/lib/godoc/images/go-logo-blue.svg"
	// fmt.Printf("URL> ")
	// fmt.Scanln(&line)
	// if strings.ToLower(line) == "exit" {
	// 	break
	// }

	req, err := grab.NewRequest("./downloads/", line)
	if err != nil {
		panic(err)
	}

	download := client.Do(req)
	bar := initBar(download)

	wg.Add(1)
	go func() {
		defer wg.Done()
		monitorDownload(download, bar)
		if err := download.Err(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	// }

	wg.Wait()
}
