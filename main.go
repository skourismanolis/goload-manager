package main

import (
	// "fmt"
	// "strings"

	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/gosuri/uiprogress"
	// "github.com/skourismanolis/goload-manager/test"
)

func initBar(name string) *uiprogress.Bar {
	bar := uiprogress.AddBar(100) // Add a new bar
	// optionally, append and prepend completion and elapsed time
	bar.Empty = '_'
	bar.Fill = '#'
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return name
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

	rand.Seed(time.Now().Unix())

	wg.Add(1)
	go func() {
		defer wg.Done()
		bar := initBar("testo.exe")

		for bar.CompletedPercent() < 100 {
			current := bar.Current()

			bar.Set(int(math.Min(float64(current)+float64(rand.Intn(5)), 100)))
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bar := initBar("picture.jpeg")

		for bar.CompletedPercent() < 100 {
			current := bar.Current()

			bar.Set(int(math.Min(float64(current)+float64(rand.Intn(5)), 100)))
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
		}

	}()

	wg.Wait()
}
