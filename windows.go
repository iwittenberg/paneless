package main

import (
	"fmt"
	"regexp"

	"github.com/JamesHovious/w32"
)

func setWindowPosition(hwnd w32.HWND, preference WindowPreference) {
	w32.MoveWindow(
		hwnd,
		int(preference.X),
		int(preference.Y),
		int(preference.Cx),
		int(preference.Cy),
		true,
	)
}

// SetWindowPositions applies the input WindowPreferences to the currently running windows
func SetWindowPositions(preference WindowPreferences) {
	w32.EnumChildWindows(0, func(hwnd w32.HWND, lparam w32.LPARAM) w32.LRESULT {
		windowTitle := w32.GetWindowText(hwnd)
		for _, pref := range preference.Preferences {
			match, _ := regexp.MatchString(pref.NameRegex, windowTitle)

			negativeMatch := false
			if len(pref.NameExlusionRegex) > 0 {
				negativeMatch, _ = regexp.MatchString(pref.NameExlusionRegex, windowTitle)
			}

			if match && !negativeMatch {
				setWindowPosition(hwnd, pref)
			}
		}
		return 1
	}, 0)
}

// GetCurrentWindowPositions returns a pointer to a WindowPreferences struct representing the current layout of all windows.
func GetCurrentWindowPositions() *WindowPreferences {
	preferences := new(WindowPreferences)
	preferences.Name = "current"
	preferences.Preferences = make([]WindowPreference, 0, 10)
	w32.EnumChildWindows(0, func(hwnd w32.HWND, lparam w32.LPARAM) w32.LRESULT {
		windowTitle := w32.GetWindowText(hwnd)
		if len(windowTitle) > 0 && windowTitle != "Default IME" && windowTitle != "MSCTFIME UI" {
			fmt.Println(windowTitle)

			windowRect := w32.GetWindowRect(hwnd)

			preference := WindowPreference{
				NameRegex:         windowTitle,
				NameExlusionRegex: "",
				X:                 windowRect.Left,
				Y:                 windowRect.Top,
				Cx:                windowRect.Right - windowRect.Left,
				Cy:                windowRect.Bottom - windowRect.Top,
			}

			preferences.Preferences = append(preferences.Preferences, preference)
		}

		return 1
	}, 0)

	return preferences
}
