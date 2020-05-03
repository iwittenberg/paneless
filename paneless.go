package main

import (
	"flag"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/iwittenberg/paneless/icon"
)

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	prefsLoc := flag.String("preferences", "preferences.json", "Preferences file path")
	snapshotLoc := flag.String("snapshot", "snapshot.json", "Snapshot file path")
	flag.Parse()

	arrangements, err := NewFromFile(*prefsLoc)
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
				openFile(*prefsLoc)
			case <-current.ClickedCh:
				s := Arrangements{*GetCurrentWindowPositions()}
				s.ToJSONFile(*snapshotLoc)
				openFile(*snapshotLoc)
			case <-quit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}
