package main

import "io/ioutil"
import "encoding/json"

// WindowPreferences represent a named preconfigured layout containing a []WindowPreference.
type WindowPreferences struct {
	Preferences []WindowPreference
	Name        string
}

// WindowPreference represents a singular application window and it's desired position keyed by a Name Regex.  An exclusion name regex can be used for more fine-grained filtering
type WindowPreference struct {
	NameRegex         string
	NameExlusionRegex string
	X                 int32
	Y                 int32
	Cx                int32
	Cy                int32
}

// WindowPreferencesList list of window prefs
type WindowPreferencesList []WindowPreferences

// ToFile writes a WindowPreferencesList to the input file name in json, creating it if it doesn't exist.
func (w *WindowPreferencesList) ToFile(filename string) error {
	data, err := json.MarshalIndent(w, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// FromFile reads the input file from json and attempts to unmarshal the contents into itself.
func (w *WindowPreferencesList) FromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, w)
	if err != nil {
		return err
	}

	return nil
}
