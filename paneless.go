package main

import (
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/iwittenberg/paneless/icon"
)

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	arrangements, err := NewFromFile("preferences.json")
	if err != nil {
		log.Fatal("Couldnt read from file", err)
	}

	systray.SetIcon(icon.Data)
	systray.SetTitle("Rearrange")
	systray.SetTooltip("Better Rearrangement Tool")

	for _, a := range *arrangements {
		item := systray.AddMenuItem(a.Name, "Rearrange the windows according to this preference")
		go func(a Arrangement) {
			for {
				<-item.ClickedCh
				a.Apply()
			}
		}(a)
	}

	systray.AddSeparator()
	current := systray.AddMenuItem("Get Snapshop", "Get a snapshop of the current positions of all windows")
	file := systray.AddMenuItem("Open Preferences", "Open the current preferences file")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-file.ClickedCh:
				openFile("preferences.json")
			case <-current.ClickedCh:
				s := Arrangements{*GetCurrentWindowPositions()}
				s.ToJSONFile("snapshot.json")
				openFile("snapshot.json")
			case <-quit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}
