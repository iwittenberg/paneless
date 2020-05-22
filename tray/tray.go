package tray

import (
	"github.com/getlantern/systray"
	"github.com/iwittenberg/paneless/arrangements"
	"github.com/iwittenberg/paneless/handler"
	"github.com/iwittenberg/paneless/icon"
	"os"
	"sync"
)

var (
	i    *Tray
	once sync.Once
)

// Try represents a system tray menu for an application.
type Tray struct {
	handler      handler.Handler
	arrangements *arrangements.Arrangements
	prefsLoc     *string
	snapshotLoc  *string
}

// GetInstance returns a singleton instance of a Tray.
func GetInstance(h handler.Handler, a *arrangements.Arrangements, p *string, s *string) (t *Tray) {
	if i == nil {
		once.Do(
			func() {
				i = &Tray{h, a, p, s}
			})
	}
	return i
}

// Init creates the actual system tray resource and begins its main execution loop.
func (t *Tray) Init() {
	systray.Run(onReady, nil)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Rearrange")
	systray.SetTooltip("Better Rearrangement Tool")

	for _, a := range *i.arrangements {
		item := systray.AddMenuItem(a.Name, "Rearrange the windows according to this preference")
		go func(a arrangements.Arrangement) {
			for {
				<-item.ClickedCh
				i.handler.Apply(&a)
			}
		}(a)
	}

	systray.AddSeparator()
	current := systray.AddMenuItem("Get Snapshot", "Get a snapshop of the current positions of all windows")
	file := systray.AddMenuItem("Open Preferences", "Open the current preferences file")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-file.ClickedCh:
				i.handler.OpenFile(*i.prefsLoc)
			case <-current.ClickedCh:
				s := arrangements.Arrangements{*i.handler.GetCurrentWindowPositions()}
				s.ToJSONFile(*i.snapshotLoc)
				i.handler.OpenFile(*i.snapshotLoc)
			case <-quit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}
