package main

import (
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/iwittenberg/paneless/icon"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	var preferences WindowPreferencesList
	if err := preferences.FromFile("preferences.json"); err != nil {
		log.Fatalf("failed to read preferences from preferences.json: %s", err)
	}

	systray.SetIcon(icon.Data)
	systray.SetTitle("Rearrange")
	systray.SetTooltip("Better Rearrangement Tool")

	for _, p := range preferences {
		go func() {
			for {
				i := systray.AddMenuItem(p.Name, "Rearrange the windows according to this preference")
				<-i.ClickedCh
				SetWindowPositions(p)
			}
		}
	}

	systray.AddSeparator()
	current := systray.AddMenuItem("Get Snapshop", "Get a snapshop of the current positions of all windows")
	file := systray.AddMenuItem("Open Preferences", "Open the current preferences file")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	// is this goroutine necessary? might be due to systray, but generally a go service will exit
	// if the main goroutine ends, even if it's spawned child routines
	go func() {
		for {
			select {
			case <-file.ClickedCh:
				open.Start("preferences.json")
			case <-current.ClickedCh:
				s := WindowPreferencesList{*GetCurrentWindowPositions()}
				s.ToFile("snapshot.json")
				open.Start("snapshot.json")
			case <-quit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}
