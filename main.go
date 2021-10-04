package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/gosuri/uiprogress"
	"github.com/inhies/go-bytesize"
	"github.com/rivo/tview"
	"github.com/skourismanolis/goload-manager/progress"
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
		progress := fmt.Sprintf("%.1f", download.Progress()*100)
		rate := bytesize.New(download.BytesPerSecond()).String() + "/s"

		eta := ""
		if time.Until(download.ETA()) >= 0 {
			eta = time.Until(download.ETA()).Round(time.Second).String()
		}

		return progress + "% " + rate + " " + eta
	})

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

func updateTable(app *tview.Application, table *tview.Table) {
	t := time.NewTicker(time.Millisecond * 150)
	defer t.Stop()
	prog := 0.0

Loop:
	for {
		select {
		case <-t.C:
			app.QueueUpdateDraw(func() {
				table.SetCellSimple(1, 1, progress.GetBar(prog, 50))
				table.SetCellSimple(1, 2, fmt.Sprintf("%0.0f%% ", prog*100))
			})
			prog += 0.01
			if prog > 1 {
				break Loop
			}
		}
	}
}

// func init() {
// }

func main() {
	table := tview.NewTable().SetBorders(true)
	table.SetBorder(true).SetTitle(" [::b]Goload [::-] Manager ")

	table.SetCellSimple(0, 0, "[::b]Filename")
	table.SetCellSimple(0, 1, "[::b]Progress")
	table.SetCellSimple(0, 2, "[::b]% Done")
	table.SetCellSimple(0, 3, "[::b]ETA")
	table.SetCellSimple(1, 0, "Giorgio_By_Moroder.mp3")
	table.SetCellSimple(1, 1, "[=======================================>----------]")
	table.SetCellSimple(1, 2, "40%")
	table.SetCellSimple(1, 3, "5s")

	// flex := tview.NewFlex()
	// flex.SetBorder(true).SetTitle("[red]Hello, [::ub]world!")
	// flex.AddItem(table, 0, 1, true)

	app := tview.NewApplication()
	go updateTable(app, table)
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}

	panic("boing")
	uiprogress.Start() // start rendering
	var wg sync.WaitGroup

	client := grab.NewClient()

	// downloads := make([]int, 0)

	for {
		var line string
		// line = "https://sabnzbd.org/tests/internetspeed/50MB.bin"
		// line = "https://golang.org/lib/godoc/images/go-logo-blue.svg"
		fmt.Printf("URL> ")
		fmt.Scanln(&line)
		if strings.ToLower(line) == "exit" {
			break
		}

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

	}

	wg.Wait()
}
