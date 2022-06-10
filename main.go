package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
)

//go:embed load.bmp
var loadingButtonImg []byte

//go:embed open.bmp
var openButtonImg []byte

func main() {
	dir := os.TempDir()
	err := ioutil.WriteFile(filepath.Join(dir, "load.bmp"), loadingButtonImg, 0600)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filepath.Join(dir, "open.bmp"), openButtonImg, 0600)
	if err != nil {
		log.Fatal(err)
	}
	open := bitmap.Open(filepath.Join(dir, "open.bmp"))
	loading := bitmap.Open(filepath.Join(dir, "load.bmp"))
	for {
		time.Sleep(time.Second)
		screen := robotgo.CaptureScreen()
		fx, fy := bitmap.Find(open, screen)
		fmt.Println("FindBitmap------ ", fx, fy)
		if fx >= 0 && fy >= 0 {
			lx, ly := bitmap.Find(loading, screen)
			if lx >= 0 || ly >= 0 {
				continue
			}

			robotgo.Move(fx+20, fy+20)
			robotgo.Click()
			robotgo.Move(fx+60, fy-40)
			robotgo.Click()
		}
	}
}
