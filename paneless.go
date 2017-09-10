package main

import (
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/iwittenberg/paneless/icon"
)

func main() {
	systray.Run(onReady)
}

func onReady() {
	preferences, err := FromJSONFile("preferences.json")
	if err != nil {
		log.Fatal("Couldnt read from file", err)
	}

	systray.SetIcon(icon.Data)
	systray.SetTitle("Rearrange")
	systray.SetTooltip("Better Rearrangement Tool")

	menuItems := make([]menuItemToName, len(*preferences))
	for i, preference := range *preferences {
		menuItems[i] = menuItemToName{
			systray.AddMenuItem(preference.Name, "Rearrange the windows according to this preference"),
			preference.Name,
		}
	}
	current := systray.AddMenuItem("Get Current", "Get the current positions of all windows")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	for _, item := range menuItems {
		go func(menuItem menuItemToName) {
			for {
				<-menuItem.item.ClickedCh

				var preference WindowPreferences
				for _, prefs := range *preferences {
					if menuItem.name == prefs.Name {
						preference = prefs
					}
				}

				SetWindowPositions(preference)
			}
		}(item)
	}

	go func() {
		select {
		case <-current.ClickedCh:
			currentPreferences := []WindowPreferences{*GetCurrentWindowPositions()}
			ToJSONFile(&currentPreferences, "currentPreferences.json")
		case <-quit.ClickedCh:
			systray.Quit()
			os.Exit(0)
		}
	}()
}

type menuItemToName struct {
	item *systray.MenuItem
	name string
}
