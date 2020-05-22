// +build windows

package handler

import (
	"github.com/JamesHovious/w32"
	"github.com/iwittenberg/paneless/arrangements"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	cmd      = "url.dll,FileProtocolHandler"
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

type Handler struct {}

// OpenFile uses the default application to open the specified file.
func OpenFile(file string) {
	_ = exec.Command(runDll32, cmd, file).Start()
}

// Apply repositions the currently running windows according to the input arrangements
func Apply(a *arrangements.Arrangement) {
	w32.EnumChildWindows(0, func(hwnd w32.HWND, lparam w32.LPARAM) w32.LRESULT {
		title := w32.GetWindowText(hwnd)
		for _, w := range a.Windows {
			match, _ := regexp.MatchString(w.NameRegex, title)

			negativeMatch := false
			if len(w.NameExclusionRegex) > 0 {
				negativeMatch, _ = regexp.MatchString(w.NameExclusionRegex, title)
			}

			if match && !negativeMatch {
				w32.MoveWindow(
					hwnd,
					int(w.X),
					int(w.Y),
					int(w.Cx),
					int(w.Cy),
					true,
				)
			}
		}
		return 1
	}, 0)
}

// GetCurrentWindowPositions returns a pointer to a WindowPreferences struct representing the current layout of all windows.
func GetCurrentWindowPositions() *arrangements.Arrangement {
	a := new(arrangements.Arrangement)
	a.Name = "current"
	a.Windows = make([]arrangements.Window, 0, 10)
	w32.EnumChildWindows(0, func(hwnd w32.HWND, lparam w32.LPARAM) w32.LRESULT {
		title := w32.GetWindowText(hwnd)
		if len(title) > 0 && title != "Default IME" && title != "MSCTFIME UI" {
			r := w32.GetWindowRect(hwnd)

			w := arrangements.Window{
				NameRegex:          title,
				NameExclusionRegex: "",
				X:                  r.Left,
				Y:                  r.Top,
				Cx:                 r.Right - r.Left,
				Cy:                 r.Bottom - r.Top,
			}

			a.Windows = append(a.Windows, w)
		}

		return 1
	}, 0)

	return a
}

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

// RegisterHokeysAndListen assigns a hotkey for each arrangement and begins listening for their activation.
func RegisterHotkeysAndListen(as *arrangements.Arrangements) {
	go func() {
		for i, _ := range *as {
			if i >= 9 {
				break
			}
			id := i + 1
			err := w32.RegisterHotKey(0, id, ModCtrl + ModAlt, uint(strconv.Itoa(id)[0]))
			if err != nil {
				log.Println("Failed to register hotkey", err)
			}
		}

		for {
			var msg = &w32.MSG{}
			w32.GetMessage(msg, 0, 0, 0)
			// Registered id is in the WPARAM field:
			if id := msg.WParam; id != 0 {
				Apply(&(*as)[id - 1])
			}
		}
	}()
}