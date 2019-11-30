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

// ToJSONFile writes a *[]WindowPreferences to the input file name, creating it if it doesn't exist.
func ToJSONFile(preferences *[]WindowPreferences, filename string) error {
	data, err := json.MarshalIndent(preferences, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// FromJSONFile reads the input file and attempts to unmarshal the contents into a *[]WindowPreferences.
func FromJSONFile(filename string) (*[]WindowPreferences, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var windowPreferences []WindowPreferences
	err = json.Unmarshal(data, &windowPreferences)
	if err != nil {
		return nil, err
	}

	return &windowPreferences, nil
}
