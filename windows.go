package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/JamesHovious/w32"
)

var (
	cmd      = "url.dll,FileProtocolHandler"
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

func setPosition(hwnd w32.HWND, w Window) {
	w32.MoveWindow(
		hwnd,
		int(w.X),
		int(w.Y),
		int(w.Cx),
		int(w.Cy),
		true,
	)
}

func openFile(file string) {
	exec.Command(runDll32, cmd, file).Start()
}

// Apply repositions the currently running windows according to the input arrangement
func (a *Arrangement) Apply() {
	w32.EnumChildWindows(0, func(hwnd w32.HWND, lparam w32.LPARAM) w32.LRESULT {
		title := w32.GetWindowText(hwnd)
		for _, w := range a.Windows {
			match, _ := regexp.MatchString(w.NameRegex, title)

			negativeMatch := false
			if len(w.NameExclusionRegex) > 0 {
				negativeMatch, _ = regexp.MatchString(w.NameExclusionRegex, title)
			}

			if match && !negativeMatch {
				setPosition(hwnd, w)
			}
		}
		return 1
	}, 0)
}

// GetCurrentWindowPositions returns a pointer to a WindowPreferences struct representing the current layout of all windows.
func GetCurrentWindowPositions() *Arrangement {
	a := new(Arrangement)
	a.Name = "current"
	a.Windows = make([]Window, 0, 10)
	w32.EnumChildWindows(0, func(hwnd w32.HWND, lparam w32.LPARAM) w32.LRESULT {
		title := w32.GetWindowText(hwnd)
		if len(title) > 0 && title != "Default IME" && title != "MSCTFIME UI" {
			r := w32.GetWindowRect(hwnd)

			w := Window{
				NameRegex:         title,
				NameExclusionRegex: "",
				X:                 r.Left,
				Y:                 r.Top,
				Cx:                r.Right - r.Left,
				Cy:                r.Bottom - r.Top,
			}

			a.Windows = append(a.Windows, w)
		}

		return 1
	}, 0)

	return a
}
