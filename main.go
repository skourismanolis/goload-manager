package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/inhies/go-bytesize"
	"github.com/rivo/tview"
	"github.com/skourismanolis/goload-manager/progress"
)

const barSize = 50

func getFileName(download *grab.Response) string {
	tmp := strings.Split(download.Filename, "\\")
	return tmp[len(tmp)-1]
}

func getPercentage(download *grab.Response) string {
	return fmt.Sprintf("%.1f%%", download.Progress()*100)
}

func getRate(download *grab.Response) string {
	return bytesize.New(download.BytesPerSecond()).String() + "/s"
}

func getETA(download *grab.Response) string {
	if time.Until(download.ETA()) >= 0 {
		return time.Until(download.ETA()).Round(time.Second).String()
	} else {
		return ""
	}
}

func monitorDownload(app *tview.Application, table *tview.Table, index int, download *grab.Response) {
	t := time.NewTicker(time.Millisecond * 150)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			app.QueueUpdateDraw(func() {
				updateDownload(table, index, download)
			})
		case <-download.Done:
			// download is complete
			app.QueueUpdateDraw(func() {
				updateDownload(table, index, download)
			})
			break Loop
		}

	}
}

func updateDownload(table *tview.Table, index int, download *grab.Response) {
	table.SetCellSimple(index, 0, getFileName(download))
	table.SetCellSimple(index, 1, progress.GetBar(download.Progress(), barSize))
	table.SetCellSimple(index, 2, getPercentage(download))
	table.SetCellSimple(index, 3, getRate(download))
	table.SetCellSimple(index, 4, getETA(download))
}

func main() {
	table := tview.NewTable().SetBorders(true)
	table.SetBorder(false).SetTitle(" [::b]Goload [::-] Manager ")

	// headers
	table.SetCellSimple(0, 0, "[::b]Filename")
	table.SetCellSimple(0, 1, "[::b]Progress")
	table.SetCellSimple(0, 2, "[::b]% Done")
	table.SetCellSimple(0, 3, "[::b]Rate")
	table.SetCellSimple(0, 4, "[::b]ETA")

	table.SetCellSimple(1, 0, "Giorgio_By_Moroder.mp3")
	table.SetCellSimple(1, 1, progress.GetBar(0, barSize))
	table.SetCellSimple(1, 2, "40%")
	table.SetCellSimple(1, 3, "5s")

	// flex := tview.NewFlex()
	// flex.SetBorder(true).SetTitle("[red]Hello, [::ub]world!")
	// flex.AddItem(table, 0, 1, true)

	app := tview.NewApplication().SetRoot(table, true)

	var wg sync.WaitGroup

	client := grab.NewClient()

	// downloads := make([]int, 0)

	// for {
	var line string
	index := 0
	line = "https://sabnzbd.org/tests/internetspeed/50MB.bin"
	// line = "https://golang.org/lib/godoc/images/go-logo-blue.svg"

	// fmt.Printf("URL> ")
	// fmt.Scanln(&line)
	// if strings.ToLower(line) == "exit" {
	// break
	// }

	req, err := grab.NewRequest("./downloads/", line)
	if err != nil {
		panic(err)
	}

	download := client.Do(req)

	wg.Add(1)
	index++
	go func() {
		defer wg.Done()
		monitorDownload(app, table, index, download)
		if err := download.Err(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	// }

	if err := app.Run(); err != nil {
		panic(err)
	}
	wg.Wait()
}
