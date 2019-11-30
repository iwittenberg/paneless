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
	preferences, err := FromJSONFile("preferences.json")
	if err != nil {
		log.Fatal("Couldnt read from file", err)
	}

	systray.SetIcon(icon.Data)
	systray.SetTitle("Rearrange")
	systray.SetTooltip("Better Rearrangement Tool")

	for _, preference := range *preferences {
		item := systray.AddMenuItem(preference.Name, "Rearrange the windows according to this preference")

		go func(item *systray.MenuItem, preferences WindowPreferences) {
			for {
				<-item.ClickedCh
				SetWindowPositions(preferences)
			}
		}(item, preference)
	}

	systray.AddSeparator()
	current := systray.AddMenuItem("Get Snapshop", "Get a snapshop of the current positions of all windows")
	file := systray.AddMenuItem("Open Preferences", "Open the current preferences file")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-file.ClickedCh:
				open.Start("preferences.json")
			case <-current.ClickedCh:
				snapshot := []WindowPreferences{*GetCurrentWindowPositions()}
				ToJSONFile(&snapshot, "snapshot.json")
				open.Start("snapshot.json")
			case <-quit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}
